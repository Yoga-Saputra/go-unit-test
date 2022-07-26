package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestHelloVindra(t *testing.T) {
	result := HelloWorld("Yoga Saputra")
	if result != "Hello Yoga Saputra" {
		// error
		t.FailNow()
	}
	fmt.Println("tidak di eksekusi")
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

// Sub test

func TestSubTest(t *testing.T) {
	t.Run("YogaTest", func(t *testing.T) {
		result := HelloWorld("Yoga")
		require.Equal(t, "Hello Yoga", result, "Result must be 'Hello Yoga'")
	})

	t.Run("SaputraTest", func(t *testing.T) {
		result := HelloWorld("Saputra")
		require.Equal(t, "Hello Saputra", result, "Result must be 'Hello Saputra'")
	})
}
