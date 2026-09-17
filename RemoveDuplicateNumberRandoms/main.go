package main

import "fmt"

func main() {
	list := []int{41, 9, 10, 10, 20, 5, 5, 5}
	result := DeleteDuplicate(list)
	fmt.Printf("Lana resultado:", result)
}
func DeleteDuplicate(list []int) []int {
	listArray := []int{}

	if len(listArray) == 0 {
		listArray = append(listArray, list[0])
	}

	for i := 0; i < len(list); i++ {
		encontrou := false
		//fmt.Printf(" line i:%d value: %d\n", i, list[i])
		for x := 0; x < len(listArray); x++ {

			fmt.Printf(" line i:%d value: %d line x:%d value: %d\n", i, list[i], x, listArray[x])
			if int(listArray[x]) == int(list[i]) {
				encontrou = true
				fmt.Printf("diferente -------")
				//listArray = append(listArray, list[i])
				break
			}

		}
		if !encontrou {
			fmt.Printf("Valor novo encontrado (%d), inserindo...\n", list[i])
			listArray = append(listArray, list[i])
		}
	}

	return listArray
}
