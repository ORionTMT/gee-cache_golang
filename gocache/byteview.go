package gocache

// ByteView holds read-only view of bytes
type ByteView struct {
	b []byte
}

// Returns the ByteView's length
func (v ByteView) Len() int {
	return len(v.b)
}

//returns a copy of data as a byte slice
//ensures that the caller cannot modify the original data

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b)) //initialize memory space for clonedBytes with zero value
	copy(c, b)                //copy the memory in b to the new memory space
	return c
}

// returns the byte data as a string
func (v ByteView) String() string {
	return string(v.b)
}
