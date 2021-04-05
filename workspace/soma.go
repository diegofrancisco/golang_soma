package main

import(
	"fmt"
	"os"
	"strconv"
	"errors"
	"log"
	//"reflect"
)
 
func main(){

	args := os.Args[1:]
	//fmt.Println(reflect.TypeOf(args))

	soma, err := Soma(args)

	if(err != nil){
		log.Fatal(err)
	} 

	fmt.Println(soma)
}

func Soma(args []string) (int, error) {
	if(len(args) < 2){
		return -1, errors.New("Número inválido de argumentos")
	}

	//fmt.Println(args)

	valor1, err1 := strconv.Atoi(args[0])
	if(err1 != nil){
		return -1, err1
	}

	valor2, err2 := strconv.Atoi(args[1])
		if(err2 != nil){
		return -1, err2
	}

	return valor1 + valor2, nil //valor1 + valor2
}