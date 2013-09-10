package swrap

import (
	"testing"
)

func TestLen(t *testing.T) {
	sw := SWrap{0x0, 0x01, 0x02}

	if sw.Len() != 3 {
		t.Error("fail")
	}
}

func TestAdd(t *testing.T) {
	sw := SWrap{}
	sw.Add(0xa)

	if len(sw) != 1 || sw[0] != 0xa {
		t.Error("fail")
	}
}

func TestMerge(t *testing.T) {
	sw := SWrap{0x0}
	sw.Merge([]byte{0x1, 0x2, 0x3})

	if len(sw) != 4 || sw[0] != 0x0 || sw[3] != 0x3 {
		t.Error("fail")
	}
}

func TestDelete(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	sw.Delete(1)

	if len(sw) != 2 || sw[0] != 0x0 || sw[1] != 0x2 {
		t.Error("fail")
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x01, 0x02}
		sw.Delete(1)
	}
}

func TestCompare(t *testing.T) {
	sw1 := SWrap{0x0, 0x1, 0x2}
	sw2 := SWrap{0x0, 0x1, 0x2}
	sw3 := SWrap{0x0, 0x1}

	if !sw1.Compare(sw2) {
		t.Error("fail")
	}

	if sw1.Compare(sw3) {
		t.Error("fail")
	}
}

func TestPush(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	sw.Push(0x3)

	if len(sw) != 4 || sw[3] != 0x3 {
		t.Error("fail")
	}
}

func TestPop(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	x := sw.Pop()

	if x != 0x2 {
		t.Error("fail")
	}
}

func TestShift(t *testing.T) {
	sw := SWrap{0x0, 0x01}
	sw.Shift(0x02)

	if len(sw) != 3 || sw[0] != 0x2 {
		t.Error("fail")
	}
}

func TestUnShift(t *testing.T) {
	sw := SWrap{0x0, 0x01, 0x02}
	v := sw.UnShift()

	if len(sw) != 2 || v != 0x0 || sw[0] != 0x01 {
		t.Error("fail")
	}
}
