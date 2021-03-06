// Copyright 2019 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package vm

import (
	"fmt"
	"strconv"
)

// AbsºInt the absolute value of x.
func AbsºInt(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// AssignAddºIntInt adds one integer to another
func AssignAddºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr += value
	return *ptr, nil
}

// AssignBitAndºIntInt equals int &= int
func AssignBitAndºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr &= value
	return *ptr, nil
}

// AssignBitOrºIntInt equals int |= int
func AssignBitOrºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr |= value
	return *ptr, nil
}

// AssignBitXorºIntInt equals int ^= int
func AssignBitXorºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr ^= value
	return *ptr, nil
}

// AssignDivºIntInt does int /= int
func AssignDivºIntInt(ptr *int64, value int64) (int64, error) {
	if value == 0 {
		return 0, fmt.Errorf(ErrorText(ErrDivZero))
	}
	*ptr /= value
	return *ptr, nil
}

// AssignModºIntInt equals int %= int
func AssignModºIntInt(ptr *int64, value int64) (int64, error) {
	if value == 0 {
		return 0, fmt.Errorf(ErrorText(ErrDivZero))
	}
	*ptr %= value
	return *ptr, nil
}

// AssignLShiftºIntInt does int <<= int
func AssignLShiftºIntInt(ptr *int64, value int64) (int64, error) {
	if value < 0 {
		return 0, fmt.Errorf(ErrorText(ErrShift))
	}
	*ptr <<= uint64(value)
	return *ptr, nil
}

// AssignMulºIntInt equals int *= int
func AssignMulºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr *= value
	return *ptr, nil
}

// AssignRShiftºIntInt does int >>= int
func AssignRShiftºIntInt(ptr *int64, value int64) (int64, error) {
	if value < 0 {
		return 0, fmt.Errorf(ErrorText(ErrShift))
	}
	*ptr >>= uint64(value)
	return *ptr, nil
}

// AssignSubºIntInt equals int *= int
func AssignSubºIntInt(ptr *int64, value int64) (int64, error) {
	*ptr -= value
	return *ptr, nil
}

// boolºInt converts integer value to boolean 0->false, not 0 -> true
func boolºInt(val int64) int64 {
	if val != 0 {
		return 1
	}
	return 0
}

// ExpStrºInt adds string and integer in string expression
func ExpStrºInt(left string, right int64) string {
	return left + strºInt(right)
}

// floatºInt converts integer value to float
func floatºInt(val int64) float64 {
	return float64(val)
}

// IncDecºInt incriment and decriment
func IncDecºInt(ptr *int64, shift int64) (int64, error) {
	var post bool

	val := *ptr
	if (shift & 0x1) == 0 {
		post = true
		shift /= 2
	}
	*ptr += shift
	if !post {
		val += shift
	}
	return val, nil
}

// MaxºIntInt returns the maximum of two integers
func MaxºIntInt(left, right int64) int64 {
	if left < right {
		return right
	}
	return left
}

// MinºIntInt returns the minimum of two integers
func MinºIntInt(left, right int64) int64 {
	if left > right {
		return right
	}
	return left
}

// strºInt converts integer value to string
func strºInt(val int64) string {
	return strconv.FormatInt(val, 10)
}
