package main

import "testing"

func TestMain(t *testing.T) {
    expected := calc(2,2,"+") 
    if expected != 3.0 {
        t.Fatalf("Expected %.2f, but got %0.2f", 3.0, expected)
    }
}
