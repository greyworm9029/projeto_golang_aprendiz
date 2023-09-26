package main 

import (
	"fmt"
)

// Concessionária - Atendimento

// Cliente representa as informações de um cliente
type Cliente struct {
	Nome string
}

// Carro representa as informaçoes de um carro
type Carro struct {
	Nome string
	Preco float64
}

// Atendente representa as informações de um atendente
type Atendente struct {
	Nome string
}

func main() {
	// Informações dos clientes 
	rodrigo := Cliente{"Rodrigo"}
	bianca := Cliente{"Bianca"}
	paulo := Cliente{"Paulo"}

	// Informações dos carros
	audi := Carro{"Audi", 220.321}
	chevrolet := Carro{"Chevrolet", 450.728}
	ferrari := Carro{"Ferrari", 976.954}

	// Informação do atendente
	joao := Atendente{"João"}

	// Organização das informações
	clientes := []Cliente{rodrigo, bianca, paulo}
	carros := []Carro{audi, chevrolet, ferrari}

	// Resultado
	for _, cliente := range clientes {
		fmt.Printf("Olá, %s! Meu nome é %s! Aqui estão os valores dos carros:\n", cliente.Nome, joao.Nome)

		for _, carro := range carros {
			fmt.Printf("%s: R$%.3f\n", carro.Nome, carro.Preco)
		}
	}
	fmt.Println("")
}