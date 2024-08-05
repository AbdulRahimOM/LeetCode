package easy

func CanBeEqual(target []int, arr []int) bool {
    mapp:=make(map[int]int)
    for _,v:=range target{
        mapp[v]++
    }
    for _,v:=range arr{
        mapp[v]--
    }
    for _,v:=range mapp{
        if v!=0{
            return false
        }
    }
    return true
}