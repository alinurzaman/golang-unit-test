package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ali")
	if result != "Hello Ali" {
		// unit test failed
		panic("Result is not Hello Ali")
	}
}
