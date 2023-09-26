package main 

import(
	"fmt"
)

// Uso de structs

func main() {
	fmt.Println("Arquivo structs")

	 type usuario struct {
		nome string
		idade int
		endereco string
		numero int
	 }
	
	usuario1 := usuario{"Julia", 21, "Rua dos Programadores", 98}

	usuario2 := usuario{nome: "Davi", idade: 21, endereco: "Rua dos Googlers", numero: 121}

	usuario3 := usuario{nome: "Roberto", idade: 22, endereco: "Rua dos Iphones", numero: 87}

	usuarios := []usuario{usuario1, usuario2, usuario3}

	for _, usuario := range usuarios {
		fmt.Printf("Meu nome é %s, tenho %d anos de idade, moro na rua %s e o número da minha casa é %d.\n", usuario.nome, usuario.idade, usuario.endereco, usuario.numero)
	}
}