package main

import "strings"

//Covert string into slice
func stringToSlice(data string, sep string) []string {
	splittedStr := strings.Split(data, sep)
	result := []string{}

	for _, str := range splittedStr {
		trim := strings.TrimSpace(str)
		if trim != "" {
			result = append(result, trim)
		}
	}
	return result
}

//Return true if all element of filter are present in arr 
//(should I range through the filter and arr the other way round?)
func containsArr(filter []string, arr []string) bool {
	for _, item := range filter {
		for _, element := range arr {
			if item == element {
				return true
			}
		}
	}
	return false
}

//Return true if str is present in arr
func contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
