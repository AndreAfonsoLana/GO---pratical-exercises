package main

func main() {
	//fmt.Println("Hello World")
	x := "100000000000001"
	value := ""

	//fmt.Printf(" Quantos caracteres %d\n", len(x))
	//fmt.Printf("x: %d\n" + x)
	for i := len(x); i > 0; i-- {
		//fmt.Printf(" Line %v\n", i)
		value += x[i-1 : i]
		//fmt.Println(" valor dessa lina ", x[i-1:i])
	}
	//fmt.Println("value:", value)

	if value == x {
		//fmt.Println("É um palíndromo")
	} else {
		//fmt.Printf("Não é um palíndromo")
	}
}
