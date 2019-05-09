package main

import (
	"testing"
)

func TestEmptyExpr(t *testing.T) {
	expected := "Empty expression"
	result := calc("")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestEmptyRecover_1(t *testing.T) {
	expected := "Empty operation (Recoverable)"
	result := calc("1 2 3 4 + * + ")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestEmptyRecover_2(t *testing.T) {
	expected := "Empty operation (Recoverable)"
	result := calc("1 3 - 4 6 + ")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestInvalidDrop_1(t *testing.T) {
	expected := "Invalid operation (drop expression)"
	result := calc("1 2 3 * * * * = = = =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestInvalidDrop_2(t *testing.T) {
	expected := "Invalid operation (drop expression)"
	result := calc("1 + 3 4 5 * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestValidExpr_1(t *testing.T) {
	expected := "Result = 15"
	result := calc("1 2 3 4 + * + =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestValidExpr_2(t *testing.T) {
	expected := "Result = 21"
	result := calc("1 2 + 3 4 + * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}