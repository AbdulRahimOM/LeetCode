package hard

import "strings"

//Runtime: 0ms, Memory: 2.46MB (beats 100%,61% when submitted)
const (
	trillionVal = 1000000000000
	billionVal  = 1000000000
	millionVal  = 1000000
)
const (
	zero      = "Zero"
	one       = "One"
	two       = "Two"
	three     = "Three"
	four      = "Four"
	five      = "Five"
	six       = "Six"
	seven     = "Seven"
	eight     = "Eight"
	nine      = "Nine"
	ten       = "Ten"
	eleven    = "Eleven"
	twelve    = "Twelve"
	thirteen  = "Thirteen"
	fourteen  = "Fourteen"
	fifteen   = "Fifteen"
	sixteen   = "Sixteen"
	seventeen = "Seventeen"
	eighteen  = "Eighteen"
	nineteen  = "Nineteen"
	twenty    = "Twenty"
	thirty    = "Thirty"
	forty     = "Forty"
	fifty     = "Fifty"
	sixty     = "Sixty"
	seventy   = "Seventy"
	eighty    = "Eighty"
	ninety    = "Ninety"
	hundred   = " Hundred"
	thousand  = " Thousand"
	million   = " Million"
	billion   = " Billion"
	trillion  = " Trillion"
)

var tens []string = []string{"", ten, twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety}
var initials []string = []string{"", one, two, three, four, five, six, seven, eight, nine, ten,
	eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen}

func NumberToWords(num int) string {
	var word strings.Builder
	if num == 0 {
		return zero
	}

	if num == 0 {
		return zero
	}
	if num >= trillionVal {
		word.WriteString(NumberToWords(num / trillionVal))
		word.WriteString(trillion)
		num = num % trillionVal
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	if num >= billionVal {
		word.WriteString(NumberToWords(num / billionVal))
		word.WriteString(billion)
		num = num % billionVal
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	if num >= millionVal {
		word.WriteString(NumberToWords(num / millionVal))
		word.WriteString(million)
		num = num % millionVal
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	if num >= 1000 {
		word.WriteString(NumberToWords(num / 1000))
		word.WriteString(thousand)
		num = num % 1000
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	if num >= 100 {
		word.WriteString(NumberToWords(num / 100))
		word.WriteString(hundred)
		num = num % 100
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	if num >= 20 {
		word.WriteString(tens[num/10])
		num = num % 10
		if num != 0 {
			word.WriteByte(' ')
		}
	}
	word.WriteString(initials[num])
	return word.String()
}
