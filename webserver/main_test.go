package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	res := Pring1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Pring1to20()
	res++
	if res != 211 {
		t.Errorf("Tets Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("--The first test", testPrint)
	t.Run("--The second test", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins...")
	m.Run()
}
