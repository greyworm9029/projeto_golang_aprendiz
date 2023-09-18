// Em GO, todo programa faz parte de um pacote. Nós definimos isso usando a palavra-chave: 'package'
package main // Aqui, o programa pertence ao pacote 'main'    

// Esse comando faz com que torne possível a importação de arquivos do pacote 'fmt'
import (
	"fmt"
)

// Isso é uma função e tudo que for executado entre as chaves '{}' será executado
func main() {
	
	/* Esse comando é uma função do pacote 'fmt', 
	   cuja responsabilidade é imprimir o texto*/
	fmt.Print("Olá, Mundo!") 
}
