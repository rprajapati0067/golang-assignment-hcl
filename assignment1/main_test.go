package main

import "testing"

func TestBalancingEmpty(t *testing.T) {
	res1 := balancingBracket("")

	if res1 != false {
		t.Errorf("failed: expected non-empty string found empty")
	}

}

func TestBalancingPositiveCase1(t *testing.T) {
	res := balancingBracket("()[]{}(([])){[()][]}")

	if res != true {
		t.Errorf("failed: should return true")
	}

}
func TestBalancingPositiveCase2(t *testing.T) {
	res := balancingBracket("[(])")

	if res != true {
		t.Errorf("failed: should return true")
	}

}

func TestBalancingNegativeCase(t *testing.T) {
	res := balancingBracket("())[]{}")

	if res != false {
		t.Errorf("failed: should return false")
	}

}
