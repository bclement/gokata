package chop

/*
Binary search that finds index of x in arr.
Returns -1 if x is not n arr
*/
func Chop(x int, arr []int) int{
    rval := -1
    size := len(arr)
    if size > 0{
        mid := size / 2
        diff := x - arr[mid]
        switch {
        case diff == 0:
            return mid
        case diff > 0:
            ret := Chop(x, arr[mid+1:])
            if ret == -1{
                return ret
            }else{
                return ret + mid + 1
            }
        case diff < 0:
            return Chop(x, arr[:mid])
        }
    }
    return rval;
}

