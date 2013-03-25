package heap

import "sort"

/*
Heap storage interface
Needs to be sortable and have push/pop
*/
type Storage interface{
    sort.Interface
    Push(x interface{})
    Pop() interface{}
}

/*
Heap structure
Includes storage and base for tree
*/
type Heap struct{
    v Storage
    base uint
}

/*
Insert item into heap
*/
func (h *Heap) Push(item interface{}){
    h.v.Push(item)
    h.percUp()
}

/*
Remove next item from heap
*/
func (h *Heap) Pop() interface{}{
    h.v.Swap(0, h.Len()-1)
    rval := h.v.Pop()
    h.percDown()
    return rval
}

/*
Move last item in storage up to place in heap
*/
func (h *Heap) percUp(){
    child := h.Len()-1
    for child > 0{
        parent := h.getParent(child)
        if !h.v.Less(child, parent){
            return
        }
        h.v.Swap(parent, child)
        child = parent
    }
}

/*
Return parent index given child index.
Does not do bounds check
*/
func (h *Heap) getParent(child int) int{
    return (child-1)/int(h.base)
}

/*
Return index of first child given parent index.
Does not do bounds check 
*/
func (h *Heap) getChild(parent int) int{
    return (parent*int(h.base))+1
}

/*
Move first element in storage to place in heap
*/
func (h *Heap) percDown(){
    parent := 0
    for parent < h.Len(){
        child := h.smallestChild(parent)
        if child < 0 || !h.v.Less(child, parent){
            return
        }
        h.v.Swap(parent, child)
        parent = child
    }
}

/*
Return index of smallest child
If no child smaller than parent, return parent
*/
func (h *Heap) smallestChild(parent int) int{
    child := h.getChild(parent)
    last := child + int(h.base)
    smallest := parent
    for ;child < h.Len() && child < last; child+=1{
        if h.v.Less(child, smallest){
            smallest = child
        }
    }
    return smallest
}

/*
Return size of heap
*/
func (h *Heap) Len() int{
    return h.v.Len()
}

/*
Storage vector for integers
*/
type IntStorage struct{
    sort.IntSlice
}

/*
Return an initialized integer heap with base
*/
func MakeInt(base uint) *Heap{
    v := IntStorage{make([]int, 0, 10)}
    return &Heap{&v, base}
}

/*
Add element to storage
*/
func (v *IntStorage) Push(x interface{}){
    arr := v.IntSlice
    size := len(arr)
    i, ok := x.(int)
    if !ok{
        panic("Passed non int to Push")
    }
    if size == cap(arr){
        // grow cap and reset len
        arr = append(arr, make([]int, size+1)...)[:size]
    }
    v.IntSlice = append(arr, i)
}

/*
Remove last element from storage
*/
func (v *IntStorage) Pop() interface{}{
    arr := v.IntSlice
    rval, arr := arr[len(arr)-1], arr[:len(arr)-1]
    v.IntSlice = arr
    return rval
}
