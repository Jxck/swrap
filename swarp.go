package swrap

type SWrap []byte

/**
 * Unshift()v       vPush(), Add()
 *         [ 1, 2, 3 ]
 *    Shift()^     ^Pop()
 */

// create swrap from []byte
func New(a []byte) SWrap {
	return SWrap(a)
}

// get []byte from swrap
func (sw *SWrap) Bytes() []byte {
	return []byte(*sw)
}

// get length of swrap
func (sw *SWrap) Len() int {
	return len(*sw)
}

// add byte at the end of swrap
func (sw *SWrap) Add(a byte) {
	*sw = append(*sw, a)
}

// merge given []byte to swrap ([swrap.., []byte...])
func (sw *SWrap) Merge(a []byte) {
	s := *sw
	l := len(s) + len(a)
	ss := make([]byte, l, l)
	copy(ss[0:], s[:])
	copy(ss[len(s):], a)
	*sw = ss
}

// delete given index value from swrap
func (sw *SWrap) Delete(i int) {
	s := *sw
	copy(s[i:], s[i+1:])
	// s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
}

// compare given []byte with swrap
func (sw *SWrap) Compare(b []byte) bool {
	s := *sw
	if len(s) != len(b) {
		return false
	}

	for i, v := range b {
		if s[i] != v {
			return false
		}
	}
	return true
}

// add byte at the end of swrap (alias of Add())
func (sw *SWrap) Push(b byte) {
	*sw = append(*sw, b)
}

// get byte at the end of swrap
func (sw *SWrap) Pop() byte {
	s := *sw
	last := s[len(s)-1]
	s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
	return last
}

// add byte at the top of swrap
func (sw *SWrap) UnShift(b byte) {
	s := *sw
	l := len(s) + 1
	ss := make([]byte, l, l)
	ss[0] = b
	copy(ss[1:], s[:])
	*sw = ss
}

// get byte at the top of swrap
func (sw *SWrap) Shift() byte {
	s := *sw
	top := s[0]
	s[0] = 0 // GC
	*sw = s[1:]
	return top
}

// replace given index value with given byte
func (sw *SWrap) Replace(i int, b byte) {
	s := *sw
	over := i - len(s)
	if over > -1 {
		ss := make([]byte, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = b
	*sw = s
}
