package chop

import "testing"

func TestChop(t *testing.T){
    assert_equal(t, 2, Chop(1, []int{-1,0,1,2,3}))
    assert_equal(t, -1, Chop(3, []int{}))
    assert_equal(t, -1, Chop(3, []int{1}))
    assert_equal(t, 0,  Chop(1, []int{1}))
    assert_equal(t, 0,  Chop(1, []int{1, 3, 5}))
    assert_equal(t, 1,  Chop(3, []int{1, 3, 5}))
    assert_equal(t, 2,  Chop(5, []int{1, 3, 5}))
    assert_equal(t, -1, Chop(0, []int{1, 3, 5}))
    assert_equal(t, -1, Chop(2, []int{1, 3, 5}))
    assert_equal(t, -1, Chop(4, []int{1, 3, 5}))
    assert_equal(t, -1, Chop(6, []int{1, 3, 5}))
    assert_equal(t, 0,  Chop(1, []int{1, 3, 5, 7}))
    assert_equal(t, 1,  Chop(3, []int{1, 3, 5, 7}))
    assert_equal(t, 2,  Chop(5, []int{1, 3, 5, 7}))
    assert_equal(t, 3,  Chop(7, []int{1, 3, 5, 7}))
    assert_equal(t, -1, Chop(0, []int{1, 3, 5, 7}))
    assert_equal(t, -1, Chop(2, []int{1, 3, 5, 7}))
    assert_equal(t, -1, Chop(4, []int{1, 3, 5, 7}))
    assert_equal(t, -1, Chop(6, []int{1, 3, 5, 7}))
    assert_equal(t, -1, Chop(8, []int{1, 3, 5, 7}))
}

func assert_equal(t *testing.T, expected int, result int){
    if expected != result{
        t.Errorf("Expected %d got %d", expected, result)
    }
}

func assert_equal2(t *testing.T, expIndex int, expFlag bool, resIndex int,
resFlag bool){
    if expIndex != resIndex || expFlag != resFlag{
        t.Errorf("Expected %d,%t got %d,%t", expIndex,
        expFlag, resIndex, resFlag)
    }
}

func TestChop2(t *testing.T){
    index, flag := Chop2(1, []int{-1,0,1,2,3})
    assert_equal2(t, 2, true, index, flag)
    index, flag = Chop2(3, []int{})
    assert_equal2(t, 0, false, index,flag)
    index, flag = Chop2(3, []int{1})
    assert_equal2(t, 1, false, index,flag)
    index, flag = Chop2(1, []int{1})
    assert_equal2(t, 0, true, index,flag)
    index, flag = Chop2(1, []int{1, 3, 5})
    assert_equal2(t, 0, true, index,flag)
    index, flag = Chop2(3, []int{1, 3, 5})
    assert_equal2(t, 1, true, index,flag)
    index, flag = Chop2(5, []int{1, 3, 5})
    assert_equal2(t, 2, true, index,flag)
    index, flag = Chop2(0, []int{1, 3, 5})
    assert_equal2(t, 0, false, index,flag)
    index, flag = Chop2(2, []int{1, 3, 5})
    assert_equal2(t, 1, false, index,flag)
    index, flag = Chop2(4, []int{1, 3, 5})
    assert_equal2(t, 2, false, index,flag)
    index, flag = Chop2(6, []int{1, 3, 5})
    assert_equal2(t, 3, false, index,flag)
    index, flag = Chop2(1, []int{1, 3, 5, 7})
    assert_equal2(t, 0, true, index,flag)
    index, flag = Chop2(3, []int{1, 3, 5, 7})
    assert_equal2(t, 1, true,  index,flag)
    index, flag = Chop2(5, []int{1, 3, 5, 7})
    assert_equal2(t, 2, true,  index,flag)
    index, flag = Chop2(7, []int{1, 3, 5, 7})
    assert_equal2(t, 3, true,  index,flag)
    index, flag = Chop2(0, []int{1, 3, 5, 7})
    assert_equal2(t, 0, false, index,flag)
    index, flag = Chop2(2, []int{1, 3, 5, 7})
    assert_equal2(t, 1, false, index,flag)
    index, flag = Chop2(4, []int{1, 3, 5, 7})
    assert_equal2(t, 2, false, index,flag)
    index, flag = Chop2(6, []int{1, 3, 5, 7})
    assert_equal2(t, 3, false, index,flag)
    index, flag = Chop2(8, []int{1, 3, 5, 7})
    assert_equal2(t, 4, false, index,flag)
}

