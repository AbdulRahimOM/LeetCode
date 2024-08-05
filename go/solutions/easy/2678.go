package easy

func CountSeniors(details []string) int {
    var seniorCnt int8
    for _,detail:=range details{
        if detail[11:13]>"60" {
            seniorCnt++
        }
    }
    return int(seniorCnt)
}