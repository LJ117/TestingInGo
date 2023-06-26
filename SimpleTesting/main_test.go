package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	// the func will write data to the os.Stdout
	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform out test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform out test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumber(t *testing.T) {

	checkNumberTests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime number, by the definition"},
		{name: "one", input: "1", expected: "1 is not prime number, by the definition"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by the definition"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "ëĘĖÊ", expected: "Please enter a whole number!"},
	}

	for _, e := range checkNumberTests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(e.expected, res) {
			t.Errorf("%s: expected: %s ; but got %s", e.name, e.expected, res)

		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function,we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)
	defer close(doneChan)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	// 模拟标准输入, 两次+
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan

}
