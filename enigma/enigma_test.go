package enigma

import "testing"

func TestEnig(t *testing.T){
    e := New(A, I, II, III)
    e.Set("AAA")
    res, _ := e.Enc("hello Franz")
    if res != "KCUBRECDEDQ"{
        t.Errorf("problem %s", res)
    }
    e.Set("AAA")
    res, _ = e.Dec(res)
    if res != "HELLOXFRANZ"{
        t.Errorf("Problem %s", res)
    }
}

func TestRot(t *testing.T){
    r := Rotor{I,0}
    for !r.willStep(){
        r.step()
    }
    if c := r.pos + 'A'; c != 'R'{
        t.Errorf("Problem %c", c)
    }
    for i := 0; i < 26; i+=1{
        r.step()
    }
    if c := r.pos + 'A'; c != 'R'{
        t.Errorf("Problem %c", c)
    }
}

func TestSym(t *testing.T){
    e := New(A, I, II, III)
    b := byte('H')
    enc := e.codec(b)
    dec := e.codec(enc)
    if dec != b {
        t.Errorf("no sym %c", dec)
    }
}

func TestStep(t *testing.T){
    e := New(A, I, II, III)
    e.Set("ADU")
    e.step()
    assertSetting(t, "ADV", e)
    e.step()
    assertSetting(t, "AEW", e)
    e.step()
    assertSetting(t, "BFX", e)
    e.step()
    assertSetting(t, "BFY", e)
}

func assertSetting(t *testing.T, expected string, e *Machine){
    if res := e.Get(); res != expected{
        t.Errorf("Expected %s, got %s", expected, res)
    }
}


