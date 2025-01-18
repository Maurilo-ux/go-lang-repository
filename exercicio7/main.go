package main

import "fmt"

func staircase(n int32) {
    
    var i int32
    

    for i = 0; i < n;i++{
        for j := 0; j < int(n-i-1) ;j++ {
            fmt.Print(" ")
        } 
        for k := i;k >= 0; k--{
            fmt.Print("#")
        }
        fmt.Println()
    }
    
    
}
func main(){

	var n int32
	n = 5
	staircase(n)


}