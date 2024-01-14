package main

import (
	"fmt"
	"testing"
)

func TestLSys(t *testing.T) {

	rules := map[string]string{
		"A": "AB",
		"B": "B",
	}

	output := runLSystem(rules, "A", 1)
	if output != "AB" {
		t.Errorf("A->AB != %s", output)
	}
	output = runLSystem(rules, "A", 2)
	if output != "ABB" {
		t.Errorf("A->AB->ABB != %s", output)
	}
	output = runLSystem(rules, "A", 3)
	if output != "ABBB" {
		t.Errorf("A->AB->ABB->ABBB != %s", output)
	}
	fmt.Println(output)
}

func TestLSysAlgae(t *testing.T) {

	rules := map[string]string{
		"A": "AB",
		"B": "A",
	}

	output := runLSystem(rules, "A", 1)
	if output != "AB" {
		t.Errorf("A->AB != %s", output)
	}
	output = runLSystem(rules, "A", 2)
	if output != "ABA" {
		t.Errorf("A->AB->ABA != %s", output)
	}
	output = runLSystem(rules, "A", 3)
	if output != "ABAAB" {
		t.Errorf("A->AB->ABB->ABAAB != %s", output)
	}
	output = runLSystem(rules, "A", 7)
	if output != "ABAABABAABAABABAABABAABAABABAABAAB" {
		t.Errorf("... != %s", output)
	}
	fmt.Println(output)
}
