// https://leetcode.com/problems/isomorphic-strings/
// Program returns a map with keys = letters and values = array of int with position or indices of that letter in the string
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
package Strings

import (
	"fmt"
	"reflect"
)

func IsIsomorphicHelper() {
	f := fmt.Println
	f("PROGRAM 11 : Isomorphic Strings")
	f(isIsomorphic("egg", "add"))
	f()
	f()
}

func isIsomorphic(s string, t string) bool {
	f := fmt.Println
	testMap1 := returnMaps(s)
	testMap2 := returnMaps(t)
	f(s, " Map -- ", testMap1)
	f(t, " Map -- ", testMap2)
	return checkMaps(testMap1, testMap2)

}

//Returns the Map of both the strings:   map[string][]int
func returnMaps(s string) map[string][]int {
	countMap := make(map[string][]int)
	testString := []byte(s)
	length := len(s)
	tempArray := []int{}
	for i := 0; i < length-1; i++ {

		if _, found := countMap[string(testString[i])]; !found {
			tempArray = append(tempArray, i)
			for j := i + 1; j < length; j++ {
				tempStringI := string(testString[i])
				tempStringJ := string(testString[j])
				if tempStringI == tempStringJ {
					tempArray = append(tempArray, j)
				}
			}
			countMap[string(testString[i])] = tempArray[:]
		}
		tempArray = []int{}

	}
	if _, found := countMap[string(testString[length-1])]; !found {
		tempArray = append(tempArray, length-1)
		countMap[string(testString[length-1])] = tempArray
	}
	return countMap
}

//Verifies whether both the maps pass compatibility test for isomorphism
func checkMaps(testMap1, testMap2 map[string][]int) bool {
	testMap := testMap2
	var tempBool bool
	if len(testMap1) == len(testMap2) {
		for _, array1 := range testMap1 {
			tempLength1 := len(array1)
			for str, array2 := range testMap {
				tempLength2 := len(array2)
				if tempLength1 == tempLength2 {
					tempBool = reflect.DeepEqual(array1, array2)
				}
				if tempBool {
					delete(testMap, str)
					break
				}
			}
			if !tempBool {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
