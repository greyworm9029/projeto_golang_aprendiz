package main

import (
	"fmt"
)

// Variável Simples

// 'var nome_da_variavel tipo (int, string, float) = valor'

func main() {

	// Declaração da variável 'carros', bem como o seu valor 
	var carros int = 2
	

	// Exibição do Resultado
	fmt.Println(carros)

	// Adicionando um novo valor 

	var carros_novo int = 4
	fmt.Printf("Novo valor: %d", carros_novo)
}