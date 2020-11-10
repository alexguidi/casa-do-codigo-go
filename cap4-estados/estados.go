package main

import "fmt"

//Estado representa cada estado do Brasil
type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 6)

	estados["GO"] = Estado{"Goias", 6434052, "Goiania"}
	estados["PB"] = Estado{"Paraiba", 3914418, "Joao Pessoa"}
	estados["PR"] = Estado{"Parana", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergige", 22283889, "Aracaju"}

	fmt.Println(estados)
}
