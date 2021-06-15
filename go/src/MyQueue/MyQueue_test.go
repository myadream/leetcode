package MyQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {
	var obj = Constructor()
	obj.Push(1)
	obj.Push(2)

	assert.Equal(t, obj.Peek(), 1, "error")
	assert.Equal(t, obj.Pop(), 1, "error")
	assert.Equal(t, obj.Empty(), false,"error")

}