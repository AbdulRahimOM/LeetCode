package medium

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	var reducedDict []string
	for _, root := range dictionary {
		prefixFound := false
		for i, prefix := range reducedDict {
			if strings.HasPrefix(root, prefix) {
				prefixFound = true
				break
			} else if strings.HasPrefix(prefix, root) {
				reducedDict[i] = root
				prefixFound = true
				break
			}
		}
		if !prefixFound {
			reducedDict = append(reducedDict, root)
		}
	}

	// words := strings.Fields(sentence)
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for _, prefix := range reducedDict {
			if strings.HasPrefix(word, prefix) {
				words[i] = prefix
				break
			}
		}
	}
	return strings.Join(words, " ")
}
