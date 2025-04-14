package services

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	ret := AddTwoNumbers(5, 6)

	//if ret != 11 {
	//	t.Errorf("the return values doesnt match with expected output that is %d", 12)
	//}

	assert.Equal(t, 11, ret)

}
