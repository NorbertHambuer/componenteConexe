package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var n int
var matrice [][]int
var nrComponenteConexe []int

func citireMatriceAdiacenta() [][]int {
	file, err := ioutil.ReadFile("matriceAdiacenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de adiacenta.")
	}

	lines := strings.Split(string(file), "\r")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		matrix[i] = make([]int, n)

		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for j, value := range vals {
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func BFS(indexStart int) []int {
	var coadaLat, parcLat []int
	vizitatLat := make([]int, n)
	/*Marcam nodul initial si il adaugam in coada*/
	vizitatLat[indexStart] = 1
	coadaLat = append(coadaLat, indexStart)

	for len(coadaLat) != 0 {
		index := coadaLat[0]

		for i := 0; i < n; i++ {
			/*Daca exista arc cu nodul curent si vecinul curent nu a fost vizitat*/
			if matrice[index][i] == 1 && vizitatLat[i] == 0 {
				/*Adaugam nodul i in coada si il marcam ca si vizitat*/
				vizitatLat[i] = 1
				coadaLat = append(coadaLat, i)
				nrComponenteConexe[i]++
			}
		}

		/*Adaugam nodul curent in lista de rezultate*/
		parcLat = append(parcLat, index)
		/*Eliminam nodul curent din coada*/
		coadaLat = coadaLat[1:]
	}

	return parcLat
}

func getNodNevizitat(){
	for 
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nrComponenteConexe = make([]int, n)

	/*Citire nod initial de unde incep parcurgerile*/
	fmt.Println("Introduceti nodul de pornire: ")
	text, _ := reader.ReadString('\n')
	indexStart, _ := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	/*Nodurile sunt stocate in memorie de la pozitia 0*/
	indexStart--

	fmt.Println("Componente conexe:")

	compConexa := BFS(indexStart)
	fmt.Println(compConexa)
}
