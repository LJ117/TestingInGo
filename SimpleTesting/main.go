package main

import "fmt"

func main() {
	n := 6

	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definitions
	switch n {
	case 0, 1:
		return false, fmt.Sprintf("%d is not prime number, by the definition", n)

	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by the definition"
	}

	//  use the modules operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not prime number
			return false, fmt.Sprintf("%d is not a prime number because its divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}
