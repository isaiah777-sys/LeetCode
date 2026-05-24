package main

import (
	f "fmt"
)

func main() {
	f.Println("Leetcode Solutions")
	//Let's use a code to swap for testing
	a, b := 1, 2
	f.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = b, a
	f.Printf("After swap: a = %d, b = %d\n", a, b)
	//Alternative way to try the check for palindrome number
	f.Printf("Is 121 a palindrome? %v\n", isPalindrome(121))
	f.Printf("Is -121 a palindrome? %v\n", isPalindrome(-121))
	f.Printf("Is 10 a palindrome? %v\n", isPalindrome(10))

	//Print the duplicates in the list

}
