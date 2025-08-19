package main

import "fmt"

const EBULICAO_K = 373.0

func main() {
	var tempK = EBULICAO_K
	var tempC = tempK - 273.0

	fmt.Printf("A temperatura de ebuliçao da agua em kelvin = %g\n"+
		"A temperatura de ebuliçao da agua em celsius = %g\n", tempK, tempC)
}
