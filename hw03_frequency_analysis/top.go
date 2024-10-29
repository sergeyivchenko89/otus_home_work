package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
)

var validID = regexp.MustCompile(`[^ \n\t]+`)

func Top10(str string) []string {

	frequencyDict := make(map[string]uint8)
	resultDict := make(map[uint8]map[string]struct{})
	result := make([]string, 0, 10)
	var maxFrequency uint8 = 0

	i := 0
	for {
		indices := validID.FindStringIndex(str[i:])
		if indices == nil {
			break
		}
		substring := validID.FindString(str[i:])
		frequencyDict[substring]++
		if frequencyDict[substring] > maxFrequency {
			maxFrequency = frequencyDict[substring]
		}
		i += indices[1]
	}

	fmt.Println(frequencyDict)

	return result
}
