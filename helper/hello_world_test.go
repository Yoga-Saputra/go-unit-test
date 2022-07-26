package helper

import (
	"fmt"
	"testing"
)

// run unit test in go : go test

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yoga Saputra")
	if result != "Hello Yoga Saputra" {
		// error
		t.Fail()
	}

	fmt.Println("dieksekusi")
}

func TestHelloYoga(t *testing.T) {
	result := HelloWorld("Yoga Saputra")
	if result != "Hello Yoga Saputra" {
		// error
		t.FailNow()
	}
	fmt.Println("tidak di eksekusi")
}

func TestHelloKay(t *testing.T) {
	result := HelloWorld("Kay")
	if result != "Hello Kay" {
		// error
		t.Error("Harusnya 'Hi Kay'")
	}
	fmt.Println("Dieksekusi")
}

func TestHelloAqmar(t *testing.T) {
	result := HelloWorld("Aqmar")
	if result != "Hello Aqmar" {
		// error
		t.Error("Harusnya 'Hi Aqmar'")
	}
	fmt.Println("Dieksekusi")
}
