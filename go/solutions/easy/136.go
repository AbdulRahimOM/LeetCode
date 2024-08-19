package easy

// When submitted: Runtime: 14ms, Memory: 6.5MB (Beats 43%, 29%)
func SingleNumber(nums []int) int {
    present:=make(map[int]bool)
    var sum1,sum2 int
    for _,v:=range nums{
        if !present[v]{
            present[v]=true
            sum1+=v
        }else{
            sum2+=v
        }
    }
    return sum1-sum2
}

//Another varied attempt:
//When submitted: Runtime: 19ms, Memory: 6.3MB (Beats 18%, 42%)
func SingleNumber_1(nums []int) int {
    present:=make(map[int]bool)
    for _,v:=range nums{
        if !present[v]{
            present[v]=true
        }else{
            present[v]=false
        }
    }
    for k,present:=range present{
        if present{
            return k
        }
    }
    return 0
}