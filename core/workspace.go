// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package core

import (
	"math/rand"
	"reflect"
	"time"
)

// Workspace contains information of compiled source code
type Workspace struct {
	Units     []*Unit
	UnitNames map[string]int
	Objects   []IObject
	Linked    map[string]int // compiled files
	IotaID    int32
}

const (
	// DefName is the key name for stdlib
	DefName = `stdlib`

	// PubOne means the only next object is public
	PubOne = 1
	// PubAll means all objects are public
	PubAll = 2
)

// Unit is a common structure for source code
type Unit struct {
	VM        *Workspace
	Index     uint32            // Index of the Unit
	NameSpace map[string]uint32 // name space of the unit
	Included  map[uint32]bool   // false - included or true - imported units
	Lexeme    []*Lex            // The array of source code
	RunID     int               // The index of run function. Undefined (-1) - run has not yet been defined
	Name      string            // The name of the unit
	Pub       int               // Public mode
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewVM returns a new virtual machine
func NewVM() *Workspace {
	ws := Workspace{
		UnitNames: make(map[string]int),
		Units:     make([]*Unit, 0, 32),
		Objects:   make([]IObject, 0, 500),
		Linked:    make(map[string]int),
	}
	return &ws
}

// InitUnit initialize a unit structure
func (ws *Workspace) InitUnit() *Unit {
	return &Unit{
		VM:        ws,
		RunID:     Undefined,
		NameSpace: make(map[string]uint32),
		Included:  make(map[uint32]bool),
	}
}

// TypeByGoType returns the type by the go type name
func (unit *Unit) TypeByGoType(goType reflect.Type) *TypeObject {
	var name string
	switch goType.String() {
	case `int64`:
		name = `int`
	case `float64`:
		name = `float`
	case `bool`:
		name = `bool`
	case `string`:
		name = `str`
	case `int32`:
		name = `char`
	case `core.KeyValue`:
		name = `keyval`
	case `core.Range`:
		name = `range`
	case `*core.Buffer`:
		name = `buf`
	case `*core.Set`:
		name = `set`
	case `*core.Array`:
		name = `arr`
	case `*core.Map`:
		name = `map`
	default:
		return nil
	}
	if obj := unit.FindType(name); obj != nil {
		return obj.(*TypeObject)
	}
	return nil
}

// StdLib returns the pointer to Standard Library Unit
func (ws *Workspace) StdLib() *Unit {
	return ws.Unit(DefName)
}

// Unit returns the pointer to Unit by its name
func (ws *Workspace) Unit(name string) *Unit {
	return ws.Units[ws.UnitNames[name]]
}

// Run executes run block
func (ws *Workspace) Run(unitID int, cmdLine []string) (interface{}, error) {
	rt := newRunTime(ws)
	if unitID < 0 || unitID >= len(ws.Units) {
		return nil, runtimeError(rt, nil, ErrRunIndex)
	}
	unit := ws.Units[unitID]
	if unit.RunID == Undefined {
		return nil, runtimeError(rt, nil, ErrNotRun)
	}
	rt.CmdLine = cmdLine
	funcRun := ws.Objects[unit.RunID].(*FuncObject)
	errResult := rt.runCmd(&funcRun.Block)
	var result interface{}
	if errResult == nil {
		if funcRun.Block.Result != nil {
			if len(rt.Stack) == 0 {
				errResult = runtimeError(rt, nil, ErrRuntime)
			} else {
				result = rt.Stack[len(rt.Stack)-1]
			}
		}
	} else {
		rt.closeAll()
	}
	for rt.Threads.Count > 0 {
		select {
		case err := <-rt.Threads.ChError:
			rt.closeAll()
			if errResult == nil {
				errResult = err
			}
		default:
		}
	}
	rt.Threads.ChCount <- 0
	close(rt.Thread.Chan)
	close(rt.Threads.ChCount)
	close(rt.Threads.ChError)
	return result, errResult
}