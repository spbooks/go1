// addition_test.go
package main

import "testing"

func TestOnePlusOne(t *testing.T) {
	onePlusOne := 1 + 1
	if onePlusOne != 2 {
		t.Error("Expected 1 + 1 to equal 2, but got", onePlusOne)
	}
}

func TestTwoPlusTwo(t *testing.T) {
	twoPlusTwo := 2 + 2
	if twoPlusTwo != 4 {
		t.Error("Expected 2 + 2 to equal 4, but got", twoPlusTwo)
	}
}
