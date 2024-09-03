package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// Equal asserts whether expected is equal to actual or not.
// Reference: github.com/stretchr/testify/assert.Equal
func Equal(t *testing.T, expected, actual any, msgAndArgs ...interface{}) {
	if isFunction(expected) || isFunction(actual) {
		t.Errorf(fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, "cannot take func type as argument"), msgAndArgs...)
	}

	if !objectsAreEqual(expected, actual) {
		t.Errorf(fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s\n"+
			"%s", expected, actual, fmt.Sprint(msgAndArgs...)))
	}
	return
}

func isFunction(arg any) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

// objectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
