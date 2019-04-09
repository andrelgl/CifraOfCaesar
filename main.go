package main

import (
	"log"
	"strings"
)

// Caesar struct
type Caesar struct {
	Matrix map[string]string
}

// NewCaesar pao com ovo
func NewCaesar(key int) map[string]string {
	matrix := make(map[string]string)
	words := "abcdefghijklmnopqrstuvwxyz"
	i := words[len(words)-(len(words)-(key+1)):]
	w := words[:len(words)-(len(words)-(key+1))]
	cypher := i + w
	for i := 0; i < len(words); i++ {
		matrix[string(words[i])] = string(cypher[i])
	}
	return matrix
}

// Encryption for simple words
func Encryption(key int, word string) string {
	matrix := NewCaesar(key)
	output := ""
	for i := 0; i < len(word); i++ {
		output = output + matrix[string(word[i])]
	}
	return output
}

// Decrypt for simples words
func Decrypt(key int, word string) string {
	matrixO := NewCaesar(key)
	matrix := map[string]string{}
	keys := []string{}
	output := ""

	for k, v := range matrixO {
		matrix[v] = k
		keys = append(keys, v)
	}
	for _, k := range keys {
		output = output + string(matrix[k][strings.Index(k, string(word))])
	}
	return output
}

func main() {
	a := Encryption(6, "batata")
	b := Decrypt(6, "ihahah")
	log.Println(a)
	log.Println(b)

}
