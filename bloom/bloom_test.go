package bloom

import "testing"

func makeHash(i uint32) func(string) uint32 {
    return func(str string) uint32 {
        return i
    }
}

func TestSet(t *testing.T) {
    f := New(2)
    size := uint32(len(f.bits)*64)
    f.set(size, makeHash(0), "foo")
    assert(t, f.bits[0] == 1, "1 works")
    f.set(size, makeHash(63), "foo")
    assert(t, f.bits[0] == 0x8000000000000001, "63 works")
    f.set(size, makeHash(64), "foo")
    assert(t, f.bits[1] == 1, "64 works")
    assert(t, f.isSet(size, makeHash(0), "foo"), "1 is set")
    assert(t, f.isSet(size, makeHash(63), "foo"), "63 is set")
    assert(t, f.isSet(size, makeHash(64), "foo"), "64 is set")
}

func TestFilter(t *testing.T) {
    f := New(2)
    f.Add("foo")
    f.Add("bar")
    assert(t, f.Contains("foo"), "contains foo")
    assert(t, f.Contains("bar"), "contains bar")
    assert(t, !f.Contains("baz"), "not baz")
    assert(t, !f.Contains("qux"), "not qux")
}

func assert(t *testing.T, result bool, msg string) {
    if !result {
        t.Errorf(msg)
    }
}
