package main 

import (
	"testing"
)

func TestFindUniquePass (t *testing.T) {
	array := []float32 {1.0,1.0,1.0,2.0}
	var unique float32
	unique = 2.0
	result := FindUniq(array)
	if result != unique {
				t.Fatalf("Error! Want - %f, got %f", unique, result)
			}
}

func TestFindUnique1 (t *testing.T) {
	array := []float32 {0.0,0.0,1.0,0.0}
	var unique float32
	unique = 1.0
	result := FindUniq(array)
	if result != unique {
				t.Fatalf("Error! Want - %f, got %f", unique, result)
			}
}

func TestFindUnique2 (t *testing.T) {
	array := []float32 {0.0,0.0,1.0,1.0, 2.0, 2.0, 3.3}
	var unique float32
	unique = 3.3
	result := FindUniq(array)
	if result != unique {
				t.Fatalf("Error! Want - %f, got %f", unique, result)
			}
}