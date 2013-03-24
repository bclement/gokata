package enigma

import "testing"

func TestEnig(t *testing.T){
    e := New(A, I, II, III)
    e.Set("AAA")
    res, _ := e.Enc("hello")
    if res != "XLEEF"{
        t.Errorf("problem %s", res)
    }
    res, _ = e.Dec(res)
    if res != "HELLO"{
        t.Errorf("Problem %s", res)
    }
}
