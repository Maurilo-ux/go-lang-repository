package main

import "fmt"

// Dados cinco inteiros positivos, encontre os valores mínimo e máximo que podem ser calculados somando exatamente quatro dos cinco inteiros. 
// Em seguida, imprima os respectivos valores mínimo e máximo como uma única linha de dois inteiros longos separados por espaço.
//Exemplo

// ARR = [1,3,5,7,9]
//A soma mínima é 1+3+5+7 = 16  e a soma máxima é 3+5+7+9 = 24. A função imprime

func miniMaxSum(arr []int32) {
    
	// Fiz um organizado crescente para facilita a soma
    tamanho := len(arr) // variavel para determinar o tamanho da Arr 

    for i := 0; i < tamanho; i++ {
        for j := 0; j < tamanho-i-1; j++ {

			//se o numero da posição j for maior que J+1, arr[j] = arr[j+1] e arr[j+1] = arr[j]
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
	// Somas precisa do int64 para numeros maiores
    var somaMax int64
    var somaMin int64
    
	// Precisa de determinar o min e o Max usando quatro dos cinco inteiros:
    for k := 0; k < tamanho -1; k++{
        somaMin += int64(arr[k]) //precisa converte o int32 pra int64
        somaMax += int64(arr[k+1])
    }
    
    fmt.Println(somaMin,somaMax)
}


func main(){

	var arr = []int32{1,7,3,5,9}
	miniMaxSum(arr)

}