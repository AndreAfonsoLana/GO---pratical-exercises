package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World Roman to integer program!")

	//roman := "MCMXCIV"
	roman := "CMXIII"
	result := romanToInt(roman)
	fmt.Println("The integer value of", roman, "is:", result)
}
func romanToInt(s string) int {
	fmt.Printf("Show the len string: %d\n", len(s))
	value := 0

	list := []struct {
		ref   string
		value int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}
	newList := []struct {
		ref   string
		value int
	}{}
	for i := 0; i < len(s); i++ {
		fmt.Printf("Numero informado: %s\n", s[i:i+1])
		for index, item := range list {
			fmt.Printf("Index: %d | Ref: %s Value: %d\n", index, item.ref, item.value)
			if s[i:i+1] == item.ref {
				newList = append(newList, struct {
					ref   string
					value int
				}{ref: item.ref, value: item.value})
			}
		}
	}
	fmt.Printf("I'm here -------------->\n")
	for i := 0; i < len(newList); i++ {
		if i+1 < len(newList) {
			if newList[i].value < newList[i+1].value {
				diff := int(newList[i+1].value) - int(newList[i].value)
				value = value + diff
				i++
			} else {
				value = value + int(newList[i].value)
			}
		} else {
			value = value + int(newList[i].value)
		}
	}
	return value

}
