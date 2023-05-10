package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

func BenchmarkHelloAli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ali")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Muhamad)",
			request:  "Muhamad",
			expected: "Hello Muhamad",
		},
		{
			name:     "HelloWorld(Ali)",
			request:  "Ali",
			expected: "Hello Ali",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Sub Test A", func(t *testing.T) {
		result := HelloWorld(("World"))
		require.Equal(t, "Hello World", result, "Result must be 'Hello World'")
		fmt.Println("TestHelloWorldRequire DONE.")
	})
	t.Run("Sub Test B", func(t *testing.T) {
		result := HelloWorld(("World"))
		assert.Equal(t, "Hello World", result, "Result must be 'Hello World'")
		fmt.Println("TestHelloWorldAssert DONE.")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before test")
	m.Run()
	fmt.Println("After test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run this test on Windows.")
	}

	result := HelloWorld(("World"))
	require.Equal(t, "Hello World", result, "Result must be 'Hello World'")
	fmt.Println("TestHelloWorldRequire DONE.")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld(("World"))
	require.Equal(t, "Hello World", result, "Result must be 'Hello World'")
	fmt.Println("TestHelloWorldRequire DONE.")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld(("World"))
	assert.Equal(t, "Hello World", result, "Result must be 'Hello World'")
	fmt.Println("TestHelloWorldAssert DONE.")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		// unit test failed
		t.Fail()
	}
	fmt.Println("TestHelloWorld done.")
}

func TestHiWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hi World" {
		// unit test failed
		t.Fail()
	}
	fmt.Println("TestHiWorld done.")
}

func TestHiAli(t *testing.T) {
	result := HelloWorld("Ali")
	if result != "Hi Ali" {
		// unit test failed
		t.FailNow()
	}
	fmt.Println("TestHiAli done.")
}

func TestHiWorldError(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hi World" {
		// unit test failed
		t.Error("Result should be Hi World")
	}
	fmt.Println("TestHiWorld done.")
}

func TestHiWorldFatal(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hi World" {
		// unit test failed
		t.Fatal("Result should be Hi World")
	}
	fmt.Println("TestHiWorld done.")
}
