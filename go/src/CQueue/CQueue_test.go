package CQueue

import (
	"testing"
)

func TestCQueue(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(3)

	if obj.DeleteHead()  == -1 {
		t.Fatal("fatal")
	}

	if obj.DeleteHead()  != -1 {
		t.Fatal("fatal")
	}

	obj = Constructor()
	if obj.DeleteHead()  != -1 {
		t.Fatal("fatal")
	}

	obj.AppendTail(5)
	obj.AppendTail(2)


	if obj.DeleteHead()  != 5 {
		t.Fatal("fatal")
	}

	if obj.DeleteHead()  != 2 {
		t.Fatal("fatal")
	}
}
/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
