package main

import "log"

// Caesar struct
type Caesar struct {
	Matrix map[string]string
}

// NewCaesar pao com ovo
func NewCaesar(key int) Caesar {
	matrix := make(map[string]string)
	words := "abcdefghijklmnopqrstuvwxyz"
	i := words[len(words)-(len(words)-(key+1)):]
	w := words[:len(words)-(len(words)-(key+1))]
	cypher := i + w
	return matrix[words] = cypher
}

func main() {
	a := NewCaesar(6)
	log.Println(a.Matrix)

}
