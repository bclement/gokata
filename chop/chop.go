package chop

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
                return ret + mid
            }
        case diff < 0:
            return Chop(x, arr[:mid])
        }
    }
    return rval;
}

