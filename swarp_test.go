package swrap

import (
	"testing"
)

func Fixture() []byte {
	return []byte{
		0x0, 0x1, 0x2, 0x3, 0x4,
		0x5, 0x6, 0x7, 0x8, 0x9,
	}
}

func TestNew(t *testing.T) {
	by := Fixture()
	sw := New(by)

	if sw[0] != 0x0 {
		t.Error("fail")
	}
}

func BenchmarkNew(b *testing.B) {
	by := Fixture()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(by)
	}
}

func TestBytes(t *testing.T) {
	sw := New(Fixture())
	var by []byte = sw.Bytes()

	if len(by) != 10 {
		t.Error("fail")
	}
}

func BenchmarkBytes(b *testing.B) {
	sw := New(Fixture())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sw.Bytes()
	}
}

func TestLen(t *testing.T) {
	sw := New(Fixture())

	if sw.Len() != 10 {
		t.Error("fail")
	}
}

func BenchmarkLen(b *testing.B) {
	sw := New(Fixture())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sw.Len()
	}
}

func TestAdd(t *testing.T) {
	sw := New(Fixture())
	sw.Add(0xFF)

	if len(sw) != 11 || sw[10] != 0xFF {
		t.Error("fail")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Add(0xFF)
	}
}

func TestMerge(t *testing.T) {
	sw := New(Fixture())
	sw.Merge([]byte{0xA, 0xB, 0xC})

	if len(sw) != 13 || sw[0] != 0x0 || sw[12] != 0xC {
		t.Error("fail")
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Merge([]byte{0xA, 0xB, 0xC})
	}
}

func TestDelete(t *testing.T) {
	sw := New(Fixture())
	sw.Delete(1)

	if len(sw) != 9 || sw[0] != 0x0 || sw[1] != 0x2 {
		t.Error("fail")
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Delete(1)
	}
}

func TestCompare(t *testing.T) {
	sw := New(Fixture())
	by1 := Fixture()
	by2 := []byte{0x0, 0x1, 0x2}

	if !sw.Compare(by1) {
		t.Error("fail")
	}

	if sw.Compare(by2) {
		t.Error("fail")
	}
}

func BenchmarkCompare(b *testing.B) {
	sw := New(Fixture())
	by := []byte{
		0x0, 0x1, 0x2, 0x3, 0x4,
		0x5, 0x6, 0x7, 0x8, 0x0, // not 0x9
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sw.Compare(by)
	}
}

func TestPush(t *testing.T) {
	sw := New(Fixture())
	sw.Push(0xFF)

	if len(sw) != 11 || sw[10] != 0xFF {
		t.Error("fail")
	}
}

func BenchmarkPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Push(0xFF)
	}
}

func TestPop(t *testing.T) {
	sw := New(Fixture())
	x := sw.Pop()

	if x != 0x9 {
		t.Error("fail")
	}
}

func BenchmarkPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Pop()
	}
}

func TestShift(t *testing.T) {
	sw := New(Fixture())
	sw.Shift(0xFF)

	if len(sw) != 11 || sw[0] != 0xFF || sw[10] != 0x9 {
		t.Error("fail")
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Shift(0xFF)
	}
}

func TestUnShift(t *testing.T) {
	sw := New(Fixture())
	v := sw.UnShift()

	if len(sw) != 9 || v != 0x0 || sw[0] != 0x1 || sw[8] != 0x9 {
		t.Error("fail")
	}
}

func BenchmarkUnShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.UnShift()
	}
}

func TestReplace(t *testing.T) {
	sw := New(Fixture())

	sw.Replace(5, 0xFF)
	if len(sw) != 10 || sw[5] != 0xFF {
		t.Error("fail")
	}

	sw.Replace(100, 0xFF)
	if len(sw) != 101 || sw[100] != 0xFF {
		t.Error("fail")
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Replace(100, 0xFF)
	}
}

func TestFor(t *testing.T) {
	by := Fixture()
	sw := New(by)

	for i, v := range sw {
		if sw[i] != v {
			t.Error("fail")
		}
	}
}
