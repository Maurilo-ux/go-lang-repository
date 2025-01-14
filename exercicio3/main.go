package main

//faça um função que receba um ano e retorne se é bisexto

import "fmt"

func anoBisexto(x int) bool{

    var resultd bool
    x = 1964
    if x%4 == 0 && x%100 != 0 && x%400 != 1{
        fmt.Println("é bissexto ")
        resultd = true
    } else {
        fmt.Println("não é bissexto ")
        resultd = false
    }
    
    return resultd
}

func main(){
    var x int
    rsl := anoBisexto(x)
    fmt.Println(rsl)
}