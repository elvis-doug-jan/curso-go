package main

import "fmt"

func main() {
	// O valor informado é para definir o tamanho o slice
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// O segundo número é para informar o tamanho do array da qual está sendo referenciado o slice
	s = make([]int, 10, 20)
	// len é o tamanho do slice e cap (capacity) é a capacidade total do array que o slice está referenciando
	fmt.Println(s, len(s), cap(s))

	// append é como o array.push do JS
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// aqui foi adicionado mais um item no array, e mesmo assim não gerou excessão, isso porque o slice vai trabalhando com array dinamico
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
