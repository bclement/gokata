package chop

import "testing"

func TestChop(t *testing.T){
    const in, out = 1, 2
    if x := Chop(in); x != out {
        t.Errorf("chop is busted")
    }
}
