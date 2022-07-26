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

// menghitung kecepatan code
func BenchmarkCategoryService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Yoga Saputra")
	}
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

func BenchmarkSubTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("YogaTest", func(b *testing.B) {
			HelloWorld("Yoga")
		})

		b.Run("SaputraTest", func(b *testing.B) {
			HelloWorld("Saputra")
		})
	}
}
