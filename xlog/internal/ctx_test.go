package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/goqianjin/common-libs/xlog/internal/assert"
)

func TestContextAttributes(t *testing.T) {
	parent := context.Background()

	// Set And Get Attribute
	intV := int64(123)
	ctx := PutContextAttribute(parent, "KeyInt", intV)
	assert.Equal(t, intV, GetContextAttribute(ctx, "KeyInt"))
	assert.Equal(t, nil, GetContextAttribute(parent, "KeyInt"))

	// Set And Get Attribute
	stringV := "stringValue"
	ctx = PutContextAttribute(ctx, "KeyString", stringV)
	assert.Equal(t, stringV, GetContextAttribute(ctx, "KeyString"))
	assert.Equal(t, nil, GetContextAttribute(parent, "KeyString"))

	// Get Attributes
	assert.Equal(t, 2, len(GetContextAttributes(ctx)))
	assert.Equal(t, 0, len(GetContextAttributes(parent)))
}

func TestName(t *testing.T) {
	m := make([]string, 0)
	m = append(m, "1")
	m = append(m, "2")
	m = append(m, "3")
	updateMap(m)
	//var ks []string
	//for k, _ := range m {
	//	ks = append(ks, k)
	//}
	fmt.Println(m)

}

func updateMap(m []string) {
	m = append(m, "A")
	m[1] = "B"
}
