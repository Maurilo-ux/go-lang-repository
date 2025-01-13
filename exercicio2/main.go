package main

// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

import (
	"fmt"
)

// len(x) - 1 dá o índice do último caractere da string.
// O loop começa no último índice (len(x) - 1) e continua até i >= 0.
// A cada iteração, o índice é decrementado com i--.

// func isPalindrome(x int) bool {

// }

func palindrome(x string) (bool) {

	x = "212212"
	var array1 []rune
	var array2 []rune


	for _, letra:= range(x){
		array1 = append(array1, rune(letra))
		
	}
		fmt.Println(array1)

	for i := len(x) -1; i >= 0; i--{
		array2 = append(array2, rune(x[i]))
	}
		fmt.Println(array2)

	
	if len(array1) != len(array2){
		return false
	}
	for j,v := range array1{
		if v != array2[j]{
			return false
		}
	}

	return true
}



func main(){
	
	var x string
	result := palindrome(x)
	fmt.Println("o palindrome é:",result)

}
