package main

import "fmt"

func Fizzbuzz(n int32) {
    // Write your code here
    n = 15
    var i int32

    for i = 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
    }
		
}

func main(){

	var n int32
    
	Fizzbuzz(n)
}