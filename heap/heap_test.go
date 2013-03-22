package heap

import "testing"
import "math/rand"

func TestHeap(t *testing.T){
    h := MakeInt(2)
    for i := 1; i < 200; i+=1{
        for n := 0; n < i; n+=1{
            h.Push(rand.Intn(100))
        }
        if l := h.Len(); l != i {
            t.Errorf("busted %d, %h", l, h)
        }
        for last,_ := h.Pop().(int); h.Len() > 0;{
            next,_ := h.Pop().(int)
            if next > last{
                t.Errorf("not sorted")
            }
        }
    }
}

