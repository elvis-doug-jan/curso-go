package main

// Ao colocar uma letra na frente do nome do pacote, criamos um alias para aquele import
// Se passar um _ (underline) na frente do import, quer dizer que esse importe não é utilizado e não deve ser removido
// pelo compilador

import (
	_ "context"
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	// forma reduzida de criar var
	area := PI * m.Pow(raio, 2)

	fmt.Printf("Área: %v \n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	g, h, i := 2, false, "teste"

	fmt.Printf("Com declaração %v, %v \n", e, f)

	fmt.Printf("Abreviado, %v, %v, %v", g, h, i)
}
