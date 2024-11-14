package main

import "fmt"

func main () {
	fib := fibgen()
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ") //3. jawaban dari step 3 nantinya akan melakukan perulangan disini
	}
}

func fibgen() func () int { // 1. tempat dimana perhitungan fibonaci dibuat
	f1 := 0
	f2 := 1
	return func() int {
		f2,f1 = (f1+f2), f2 
		return f1 // 2. jawaban dari f1 + f2 nantinya disimpan di f2 sebelum di kembalikan ke f1
	}
}