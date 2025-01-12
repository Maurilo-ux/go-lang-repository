package main

// Escreva um algoritmo que armazene o valor 10 em uma variável A e o valor 20 em uma variável B. 
// A seguir (utilizando apenas atribuições entre variáveis) troque os seus conteúdos fazendo com que o 
// valor que está em A passe para B e vice-versa. Ao final, escrever os valores que ficaram armazenados 
// nas variáveis.


import "fmt"

//observação: de ideia inicial eu pensei em criar uma 3 variavel para salvar o valores. mas pequisei e vi que posso simplesmente declarar 2 variavel e seus valores

func trocaValores(a, b int)(int,int){

	//ideia inicial 
	// a = c 
	// b = a 
	// c = b

	// A = B e B = A.
	a,b = b,a
	return a,b

}

func main(){
	a := 10
	b := 20

	a, b = trocaValores(a,b)
	fmt.Println(a,b)
}