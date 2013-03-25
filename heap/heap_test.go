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
            if next < last{
                t.Errorf("not sorted")
            }
            last = next
        }
    }
}

type Reverse struct{
   IntStorage
}

func (r *Reverse) Less(i, j int) bool{
    return r.IntStorage.Less(j, i)
}

func TestReverse(t *testing.T){
    v := Reverse{IntStorage{make([]int, 0, 10)}}
    h := &Heap{&v, 2}
    h.Push(2)
    h.Push(3)
    h.Push(1)
    for last,_ := h.Pop().(int); h.Len() > 0;{
        next,_ := h.Pop().(int)
        if next > last{
            t.Errorf("not reversed")
        }
        last = next
    }
}
