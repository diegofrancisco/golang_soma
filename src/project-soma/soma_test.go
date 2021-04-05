package main

import (
    "testing"
    //"regexp"
)

func TestSoma_Base(t *testing.T) {
	args := []string{"3", "4"}
    soma, err := Soma(args)
	soma_check := 3 + 4
    if soma != soma_check || err != nil {
        t.Fatalf(`Soma([3, 4]) = %v, %v, soma inválida`, soma, err)
    }
}

func TestSoma_Args(t *testing.T) {
	args := []string{"3"}
    soma, err := Soma(args)
    if err == nil {
        t.Fatalf(`Soma([3]) = %v, %v, aceitando numero inválido de argumentos`, soma, err)
    }
}

func TestSoma_String(t *testing.T) {
	args := []string{"2", "a"}
    soma, err := Soma(args)
    if err == nil {
        t.Fatalf(`Soma([3, a]) = %v, %v, aceitando argumentos não numericos`, soma, err)
    }
}