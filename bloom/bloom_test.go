package bloom

import "os"
import "io"
import "strings"
import "bufio"
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

func TestGerman(t *testing.T) {
    f := New(512)
    load(t, f, "words.txt")
    file, err := os.Open("german.txt")
	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                t.Errorf("%s", err)
            }
            break;
        }
        if f.Contains(strings.TrimSpace(line)) {
            t.Errorf("got %s", line)
        }
    }
}

func load(t *testing.T, f *Filter, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                t.Errorf("%s", err)
            }
            break;
        }
        f.Add(strings.TrimSpace(line))
    }
}

func assert(t *testing.T, result bool, msg string) {
    if !result {
        t.Errorf(msg)
    }
}
