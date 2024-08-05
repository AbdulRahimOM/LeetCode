package easy

func KthDistinct(arr []string, k int) string {
    found:=make(map[string]int)
    for _,word:=range arr{
        found[word]++
    }
    for _,word:=range arr{
        if found[word]==1{
            k--
            if k==0{
                return word
            }
        }
    }
    return ""
}