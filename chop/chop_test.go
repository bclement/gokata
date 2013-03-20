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
