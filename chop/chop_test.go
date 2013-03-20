package chop

import "testing"

func TestChop(t *testing.T){
    assert(t, 2, Chop(1, []int{-1,0,1,2,3}))
}

func assert(t *testing.T, expected int, result int){
    if expected != result{
        t.Errorf("Expected %d got %d", expected, result)
    }
}
