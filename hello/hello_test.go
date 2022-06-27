package hello_test

import (
	"golearn/hello"
	"testing"
)

func TestHello(t *testing.T) {
	if hello.Greet() != "Hi there" {
		t.Fatal("Wrong greeting")
	}
}
