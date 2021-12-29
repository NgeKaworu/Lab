package generate

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []int{5}
	for _, v := range tests {
		log.Println(generate(v))
	}
}
