# Conversor de Temperatura em Go

Este é um projeto simples desenvolvido para a DIO (Digital Innovation One) que converte a temperatura de ebulição da água de Kelvin para Celsius utilizando a linguagem Go.

## Descrição
O programa define a constante da temperatura de ebulição da água em Kelvin e realiza a conversão para Celsius, exibindo ambos os valores no terminal.

## Código Principal
```go
package main

import "fmt"

const EBULICAO_K = 373.0

func main() {
    var tempK = EBULICAO_K
    var tempC = tempK - 273.0

    fmt.Printf("A temperatura de ebuliçao da agua em kelvin = %g\n"+
        "A temperatura de ebuliçao da agua em celsius = %g\n", tempK, tempC)
}
```

## Como Executar
1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório ou baixe o arquivo `main.go`.
3. No terminal, navegue até a pasta do projeto e execute:

```fish
go run main.go
```

## Saída Esperada
```
A temperatura de ebuliçao da agua em kelvin = 373
A temperatura de ebuliçao da agua em celsius = 100
```

## Licença
Este projeto é apenas para fins educacionais.
