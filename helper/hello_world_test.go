package helper

import (
	"fmt"
	"testing"
)

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
