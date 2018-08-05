package intslice

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T){
	t.Run("[]int{nil},[]int{nil}", testIsEqualFunc([]int(nil),[]int(nil),true)) //compare two nil slices
	t.Run("[]int{nil},[]int{}", testIsEqualFunc([]int(nil),[]int{},false)) //compare nil slice with empty slice
	t.Run("[]int{nil},[]int{0,1,2,3}", testIsEqualFunc([]int(nil),[]int{0,1,2,3},false)) //compare nil slice with element slice
	t.Run("[]int{0,1,2,3},[]int{0,1,2,3,4,5}", testIsEqualFunc([]int{0,1,2,3},[]int{0,1,2,3,4,5},false)) //compare two slices with different length
	t.Run("[]int{0,1,2,3},[]int{0,3,2,1}", testIsEqualFunc([]int{0,1,2,3},[]int{0,3,2,1},false))// compare two slices same length different values
	t.Run("[]int{0,1,2,3},[]int{0,1,2,3}", testIsEqualFunc([]int{0,1,2,3},[]int{0,1,2,3},true)) // compare two equal slices	
}

func testIsEqualFunc(a,b []int, expected bool) func (*testing.T){
	return func (t *testing.T){
		actual := IsEqual(a,b)
		if actual != expected {
			t.Error(fmt.Sprintf("Compared %v with %v. Got %v expected %v",a,b,actual,expected)) 
		}
	}
}