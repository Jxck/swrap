package swrap

import (
	"testing"
)

func Fixture() *SWrap {
	return &SWrap{
		0x0, 0x1, 0x2, 0x3, 0x4,
		0x5, 0x6, 0x7, 0x8, 0x9,
	}
}

func TestLen(t *testing.T) {
	sw := SWrap{0x0, 0x01, 0x02}

	if sw.Len() != 3 {
		t.Error("fail")
	}
}

func BenchmarkLen(b *testing.B) {
	sw := SWrap{0x0, 0x01, 0x02}
	for i := 0; i < b.N; i++ {
		sw.Len()
	}
}

func TestAdd(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	sw.Add(0x3)

	if len(sw) != 4 || sw[3] != 0x3 {
		t.Error("fail")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x1, 0x2}
		sw.Push(0x3)
	}
}

func TestMerge(t *testing.T) {
	sw := SWrap{0x0}
	sw.Merge([]byte{0x1, 0x2, 0x3})

	if len(sw) != 4 || sw[0] != 0x0 || sw[3] != 0x3 {
		t.Error("fail")
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0}
		sw.Merge([]byte{0x1, 0x2, 0x3})
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

func BenchmarkCompare(b *testing.B) {
	sw1 := SWrap{0x0, 0x1, 0x2}
	sw2 := SWrap{0x0, 0x1, 0x2}
	for i := 0; i < b.N; i++ {
		sw1.Compare(sw2)
	}
}

func TestPush(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	sw.Push(0x3)

	if len(sw) != 4 || sw[3] != 0x3 {
		t.Error("fail")
	}
}

func BenchmarkPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x1, 0x2}
		sw.Push(0x3)
	}
}

func TestPop(t *testing.T) {
	sw := SWrap{0x0, 0x1, 0x2}
	x := sw.Pop()

	if x != 0x2 {
		t.Error("fail")
	}
}

func BenchmarkPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x1, 0x2}
		sw.Pop()
	}
}

func TestShift(t *testing.T) {
	sw := SWrap{0x0, 0x01}
	sw.Shift(0x02)

	if len(sw) != 3 || sw[0] != 0x2 {
		t.Error("fail")
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x01}
		sw.Shift(0x02)
	}
}

func TestUnShift(t *testing.T) {
	sw := SWrap{0x0, 0x01, 0x02}
	v := sw.UnShift()

	if len(sw) != 2 || v != 0x0 || sw[0] != 0x01 {
		t.Error("fail")
	}
}

func BenchmarkUnShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := SWrap{0x0, 0x01, 0x02}
		sw.UnShift()
	}
}
