package main

import (
	"fmt"
	"sort"
)

/*
CÓDIGO AINDA INCOMPLETO

LER NOME DE FUNCIONARIOS ATRAVES DE UMA STRUCT E EXIBIR QUANTO DE ESPAÇO DO HD O FUNCIONÁRIO ESTÁ USANDO
*/

type Empresa struct {
	Nome   string
	Espaco float64
	ID     int
}

func main() {

	const tamanho = 3
	sliceNomes := []string{}

	slicePorcentagemEspaco := []float64{}

	var TempNome string
	var TempEspaco float64
	var PorcentagemEspaco float64
	//var ConsultaId int
	EspacoTotal := 1000.0
	EspacoRestante := 1000.0
	SomaEspacos := 0.0 //Soma de todos os espaços restantes

	for i := 0; i < tamanho; i++ {
		fmt.Println("Digite o nome do funcionario", i)
		fmt.Scan(&TempNome)
		fmt.Println("Digite o espaço em disco do funcionario", i)
		fmt.Scan(&TempEspaco)

		empresa := Empresa{TempNome, TempEspaco, i}
		SomaEspacos = SomaEspacos + TempEspaco

		EspacoRestante = 1000 - SomaEspacos

		sliceNomes = append(sliceNomes, empresa.Nome)

		PorcentagemEspaco = (TempEspaco / EspacoTotal) * 100

		slicePorcentagemEspaco = append(slicePorcentagemEspaco, PorcentagemEspaco)

	}

	for i := 0; i < tamanho; i++ {
		fmt.Println("ID:", i, "\t\t\t Nome:", sliceNomes[i], "\t\t\tEspaço Livre: ", slicePorcentagemEspaco[i], "%")
	}
	fmt.Println("\t\t\tEspaço Restante no HD:", EspacoRestante)

	// ===========================================================================//

	sort.Strings(sliceNomes)

}
