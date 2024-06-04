//Minimum Number of Chairs in a Waiting Room
//https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/description/
package easy

func MinimumChairs(s string) int {
    count:=0
    max:=0
    for _,v:=range s{
        if v=='E'{
            count++
        }else{
            count--
        }
        if count>max{
            max=count
        }
    }
    return max
}

