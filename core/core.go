// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package core

// CmdType is used for types of commands
type CmdType uint32

const (
	// CtValue pushes value into stack
	CtValue CmdType = iota + 1
	// CtVar pushes the value of the variable into stack
	CtVar
	// CtConst pushes the value of the constant into stack
	CtConst
	// CtStack is a stack command
	CtStack
	// CtUnary is an unary function
	CtUnary
	// CtBinary is binary function
	CtBinary
	// CtFunc is other functions
	CtFunc
)
const (
	// Undefined index
	Undefined = -1
)

const (
	// StackBlock executes function
	StackBlock = 1 + iota
	// StackReturn returns values from the function
	StackReturn
	// StackIf is the condition statement
	StackIf
	// StackWhile is the while statement
	StackWhile
	// StackAssign is an assign operator
	StackAssign
	// StackAnd is a logical AND
	StackAnd
	// StackOr is a logical OR
	StackOr
	// StackQuestion is ?(condition, exp1, exp2)
	StackQuestion
	// StackInc is ++ --
	StackIncDec
)

// Token is a lexical token.
type Token struct {
	Type   int
	Offset int
	Length int
}

// Lex contains the result of the lexical parsing
type Lex struct {
	Source []rune
	Tokens []Token
	Lines  []int // offsets of lines
}

// ICmd is an interface for stack commands
type ICmd interface {
	GetType() CmdType
	GetResult() *TypeObject
	GetObject() IObject
	GetToken() int
}

// CmdCommon is a common structure for all commands
type CmdCommon struct {
	TokenID uint32 // the index of the token in lexeme
}

// CmdValue pushes a value into stack
type CmdValue struct {
	CmdCommon
	Value  interface{}
	Result *TypeObject
}

// CmdVar pushes the value of the variable into stack
type CmdVar struct {
	CmdCommon
	Block  *CmdBlock // pointer to the block of the variable
	Index  int       // the index of the variable in the block
	LValue bool
}

// CmdConst pushes a value of the constant into stack
type CmdConst struct {
	CmdCommon
	Object IObject
}

// CmdBlock calls a stack command
type CmdBlock struct {
	CmdCommon
	Parent   *CmdBlock
	Object   IObject
	ID       uint32 // cmdType
	Vars     []*TypeObject
	ParCount int // the count of parameters
	VarNames map[string]int
	Result   *TypeObject
	Children []ICmd
}

// CmdUnary calls an unary function
type CmdUnary struct {
	CmdCommon
	Object  IObject
	Result  *TypeObject
	Operand ICmd
}

// CmdBinary calls a binary function
type CmdBinary struct {
	CmdCommon
	Object IObject
	Result *TypeObject
	Left   ICmd
	Right  ICmd
}

// CmdAnyFunc calls a function with more than 2 parameters
type CmdAnyFunc struct {
	CmdCommon
	Object   IObject
	Result   *TypeObject
	Children []ICmd
}

// GetType returns CtValue
func (cmd *CmdValue) GetType() CmdType {
	return CtValue
}

// GetResult returns result
func (cmd *CmdValue) GetResult() *TypeObject {
	return cmd.Result
}

// GetToken returns the index of the token
func (cmd *CmdValue) GetToken() int {
	return int(cmd.TokenID)
}

// GetObject returns nil
func (cmd *CmdValue) GetObject() IObject {
	return nil
}

// GetType returns CtValue
func (cmd *CmdVar) GetType() CmdType {
	return CtVar
}

// GetResult returns result
func (cmd *CmdVar) GetResult() *TypeObject {
	return cmd.Block.Vars[cmd.Index]
}

// GetToken returns teh index of the token
func (cmd *CmdVar) GetToken() int {
	return int(cmd.TokenID)
}

// GetObject returns nil
func (cmd *CmdVar) GetObject() IObject {
	return nil
}

// GetType returns CtConst
func (cmd *CmdConst) GetType() CmdType {
	return CtConst
}

// GetResult returns result
func (cmd *CmdConst) GetResult() *TypeObject {
	return cmd.Object.Result()
}

// GetToken returns the index of the token
func (cmd *CmdConst) GetToken() int {
	return int(cmd.TokenID)
}

// GetObject returns nil
func (cmd *CmdConst) GetObject() IObject {
	return cmd.Object
}

// GetType returns CtStack
func (cmd *CmdBlock) GetType() CmdType {
	return CtStack
}

// GetResult returns result
func (cmd *CmdBlock) GetResult() *TypeObject {
	return cmd.Result
}

// GetObject returns nil
func (cmd *CmdBlock) GetObject() IObject {
	return cmd.Object
}

// GetToken returns teh index of the token
func (cmd *CmdBlock) GetToken() int {
	return int(cmd.TokenID)
}

// GetType returns CtUnary
func (cmd *CmdUnary) GetType() CmdType {
	return CtUnary
}

// GetResult returns the type of the result
func (cmd *CmdUnary) GetResult() *TypeObject {
	return cmd.Result
}

// GetObject returns Object
func (cmd *CmdUnary) GetObject() IObject {
	return cmd.Object
}

// GetToken returns teh index of the token
func (cmd *CmdUnary) GetToken() int {
	return int(cmd.TokenID)
}

// GetType returns CtBinary
func (cmd *CmdBinary) GetType() CmdType {
	return CtBinary
}

// GetResult returns the type of the result
func (cmd *CmdBinary) GetResult() *TypeObject {
	return cmd.Result
}

// GetObject returns Object
func (cmd *CmdBinary) GetObject() IObject {
	return cmd.Object
}

// GetToken returns teh index of the token
func (cmd *CmdBinary) GetToken() int {
	return int(cmd.TokenID)
}

// GetType returns CtFunc
func (cmd *CmdAnyFunc) GetType() CmdType {
	return CtFunc
}

// GetResult returns the type of the result
func (cmd *CmdAnyFunc) GetResult() *TypeObject {
	return cmd.Result
}

// GetObject returns Object
func (cmd *CmdAnyFunc) GetObject() IObject {
	return cmd.Object
}

// GetToken returns teh index of the token
func (cmd *CmdAnyFunc) GetToken() int {
	return int(cmd.TokenID)
}

// LineColumn return the line and the column of the ind-th token
func (lp Lex) LineColumn(ind int) (line int, column int) {
	end := len(lp.Tokens) == ind && ind > 0
	if end {
		ind--
	}
	if len(lp.Tokens) > ind {
		for ; line < len(lp.Lines); line++ {
			if lp.Lines[line] > lp.Tokens[ind].Offset {
				break
			}
		}
		column = lp.Tokens[ind].Offset - lp.Lines[line-1] + 1
		if end {
			column += lp.Tokens[ind].Length
		}
	}
	return
}