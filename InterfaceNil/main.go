package main

import "fmt"

type DispatcherError struct{ Msg string }

// Implementa a interface error para o nosso tipo customizado
func (e *DispatcherError) Error() string { return e.Msg }

func validateDDR() error {
	// myErr tem o Valor nil, mas o Tipo já está definido (*DispatcherError)
	var myErr *DispatcherError = nil

	// A interface error retornada recebe: (Tipo: *DispatcherError, Valor: nil)
	return myErr
}
func validateDDRCorrect() error {
	// Declare como error em vez de *DispatcherError
	var err error = nil
	return err // Retorna com segurança (Tipo: nil, Valor: nil)
}

func main() {
	err := validateDDR()

	// O Go compara (Tipo: *DispatcherError, Valor: nil) com (Tipo: nil, Valor: nil)
	// Como os tipos não batem, err não é nil, e a condição dá TRUE!
	if err != nil {
		fmt.Println("Falso positivo! O código entra no erro, mesmo o valor de myErr sendo nil.")
	}
	fmt.Printf("Tipo do err: %T\n", err) // Saída: *main.DispatcherError

	err = validateDDRCorrect()

	fmt.Printf("Tipo do err: %T\n", err) // Saída: *main.DispatcherError
	if err != nil {
		fmt.Println("Falso positivo! O código entra no erro, mesmo o valor de myErr sendo nil.")
	} else {
		fmt.Println("Deu certo")
	}
}
