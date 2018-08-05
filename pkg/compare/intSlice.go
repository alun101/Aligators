package intslice

//IsEqual recieves two int slices and returns true if they are equal  
func IsEqual(a, b []int) bool {
	if a == nil && b == nil { //if slices are both 'nil' they are equal.
		return true
	}
	if a == nil || b == nil { //this comparison must come after the above. If only one slice is 'nil' they are not equal.
		return false
	}
	if len(a) != len(b) { //if slice lengths are not equal then slices are not equal.
		return false
	}
	for i := range a { //compare values in both slices one at a time.
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// a := []int(nil) is a nil slice
// var a []int is a nil slice
// a := []int{} is an empty slice
// a := make([]int,0) is an empty slice
