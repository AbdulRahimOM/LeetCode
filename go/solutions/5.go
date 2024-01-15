package solutions



func Output5(s string) any {
	return longestPalindrome(s)
	}


func longestPalindrome(s string) string {
    maxp:=1
    var start int=0
    length:=len(s)
    for i:=1;i<length;i++{
        for j:=1;i-j>=0 &&i+j<length;j++{
            if s[i+j]==s[i-j]{
                if 2*j+1>maxp{
                    maxp=2*j+1
                    start=i-j
                }
            }else{
                break
            }
        }
        for j:=1;i-j>=0 &&i+j-1<length;j++{
            if s[i+j-1]==s[i-j]{
                if 2*j>maxp{
                    maxp=2*j
                    start=i-j
                }
            }else{
                break
            }
        }
    }
    return s[start:start+maxp]
}