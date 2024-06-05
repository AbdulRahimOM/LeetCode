package easy

func CommonChars(words []string) []string {
	m := make(map[rune]byte)
	for _, letter := range words[0] {
		m[letter]=m[letter]+1
	}
	for _, word := range words[1:] {
		for key,count := range m {
			var localCount byte=0
			for _, letter := range word {
				if letter == key {
					localCount++
				}
				if localCount==count{
					break
				}
			}
			if localCount!=count{
				if localCount==0{
					delete(m,key)
				}else{
					m[key]=localCount
				}
			}
		}
	}
	result:=[]string{}
	for letter,count:=range m{
		for count>0{
			result=append(result, string(letter))
			count--
		}
	}
	return result
}
