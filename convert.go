// Package humannumbers contains utilities for presenting numbers in user friendly
package humannumbers

import "fmt"

// NumToWord will be exported if name of func is start with cap then it is exportable not the camelcase
func NumToWord(number int) (string, error) {

	if number < 0 || number >100 {
		//panic("Something went wrong") it an inbuilt error
		return "", fmt.Errorf("The number %v doesn't fall in range 0-100", number)
	}

	if number == 100 {
		return "hundred", nil
	}

	unitsWord := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
        "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
        "eighteen", "nineteen"}

		
	if number < 20 {
		return unitsWord[number],nil
	}

	tensWord := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	return tensWord[number/10]+ " " + unitsWord[number%10], nil
}