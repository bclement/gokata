package chop

/*
Binary search that finds index of x in arr.
Returns -1 if x is not n arr
*/
func Chop(x int, arr []int) int{
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
    return -1;
}

/*
Takes a value and an array that may contain that value.
If the boolean return is true, the int is the index of the value.
If the boolean return is false, the int is where the value would go.
*/
func Chop2(x int, arr []int) (int, bool){
    size := len(arr)
    if size > 0{
        mid := size / 2
        diff := x - arr[mid]
        add := 0
        switch {
        case diff == 0:
            return mid, true
        case diff > 0:
            arr = arr[mid+1:]
            add += mid + 1
        case diff < 0:
            arr = arr[:mid]
        }
        ret, flag := Chop2(x, arr)
        return ret + add, flag
    }
    return 0, false
}
