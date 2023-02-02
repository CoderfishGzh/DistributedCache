package geecache

// ByteView hold immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns the view's len
// is the Value interface
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a bytes slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
