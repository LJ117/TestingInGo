package main

import (
	"testing"
)

// just a test with nothing return
func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because its divisible by 2"},
		{"negativedPrime", -3, false, "Negative numbers are not prime, by the definition"},
		{"one", 1, false, "1 is not prime number, by the definition"},
		{"zero", 0, false, "0 is not prime number, by the definition"},
	}

	for _, e := range primeTests {
		res, msg := isPrime(e.testNum)
		if e.expected != res {
			t.Errorf("%s: expected %t but got %t", e.name, e.expected, res)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

}
