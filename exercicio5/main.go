package main

import "fmt"


func diagonalDifference(arr [][]int32) int32{

    
    var somarDiagonalDireita int32
    var somaDiagonalEsquerda int32
    tamanho := len(arr)

    for i := 0; i < tamanho; i++{
        somarDiagonalDireita += arr[i][i] 
        somaDiagonalEsquerda += arr[i][int32(tamanho) - 1 - int32(i) ]
    }

    x := int32( somarDiagonalDireita - somaDiagonalEsquerda)

    if x < 0{
		return -x
	} else{
		return x
	}
    
}

func main(){

	var arr = [][]int32{{-10,3,0,5,-4},{2,-1,0,2,-8},{9,-2,-5,6,0},{9,-7,4,8,-2},{3,7,8,-5,0}}
	// D = {-1,-8,7,-2} E = {-8,-5,7,1}

	
	result := diagonalDifference(arr)
	fmt.Println(result)

}

