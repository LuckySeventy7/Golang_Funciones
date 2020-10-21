package main

import "fmt"

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
func fibonacci(){

	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
}

func numMayor(args ...int) int{
	mayorN := 0
	for _, v := range args { 
		if(v> mayorN){
			mayorN=v
		}
	}
	return mayorN
}
func generadorImpares() func() uint {
	i := uint(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func intercambia(a *int, b *int){
	aux := a
 	a= b
	b= aux
	fmt.Println("a: ",*a,", b:",*b)
}

func main() {
	var op int
		
	fmt.Println("Funciones")
	fmt.Println("\n1. Calculadora Fibonacci \n2. Numero mayor \n3. Generador de numero impares \n4. Intercambia numeros \n5. Salir")
	fmt.Println("Ingrese opcion:")
	fmt.Scanf("%d\n",&op)

	if(op==1){
		var n int
		fmt.Print("Ingrese n:")
		fmt.Scan( &n)
		fmt.Print("Respuesta:")
		fmt.Println(fib(n))
		//fibonacci(n)
	}else if(op==2){
		var n int
		var x int
		var s []int
		fmt.Print("Ingrese cantidad de numeros que desea ingresar:")
		fmt.Scanf("%d\n", &n)
	
		for i:=0; i < n; i++ {
		//	fmt.Print(i)
			fmt.Scan( &x)
			s=append(s,x)	
		}
		fmt.Print("Numero mayor:")
		fmt.Println(numMayor(s...))


	}else if(op==3){
		var n int
		nextImpar := generadorImpares()
		fmt.Print("Ingrese n:")
		fmt.Scan( &n)
		for i:=0; i < n; i++ {
			fmt.Println(nextImpar())
		}
	}else if(op==4){
		var a int
		var b int
		fmt.Print("Ingrese a :")
		fmt.Scan( &a)
		fmt.Print("Ingrese b :")
		fmt.Scan( &b)
		intercambia(&a, &b)

	}
}