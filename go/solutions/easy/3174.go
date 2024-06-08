package easy

func ClearDigits(s string) string {
	length:=len(s)
    for i:=0;i<length;i++{
		if s[i]>='0'&&s[i]<='9'{
			s=s[:i-1]+s[i+1:]
			i-=2
			length-=2
		}
	}
	return s
}

