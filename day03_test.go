package main

import (
	"testing"
)

func TestCleanCorruptedMemory(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	output := cleanCorruptedMemory(input)

	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
		}

	}
}

func TestCalculateInstructions(t *testing.T) {
	input := "mul(2,6)"
	expected := 12
	output := calculateInstruction(input)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestCalculateFromCorruptedMemory(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	output := CalculateFromCorruptedMemory(input)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}
