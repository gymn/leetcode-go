package _0001_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a := TwoSum([]int{3, 3}, 6)
	b := []int{1, 0}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("expect [0 1] but found %v", a)
	}
}
