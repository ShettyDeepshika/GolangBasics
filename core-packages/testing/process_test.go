package main

import (
	"testing"
)
//Always start your test dunction with "Test" (Uppercase T)
func TestCheckEven( t *testing.T){
	i := 10
	expected := "YES"
	res := checkEven(i)
	if res != expected {
		t.Errorf("expected %v got %v",expected,res)
	}
}