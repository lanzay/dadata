package suggestions

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(t *testing.T)  {
	
	dada := NewDaData("b4b0dd78fe3055a1181131d3372e9e660bb3d1c1")
	
	res0, err := dada.SuggestAddresses("Москва Мира 231",10)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	
	fmt.Println("[Unrestricted_value]",res0[0].Unrestricted_value)
	res1, _ := dada.SuggestAddresses(res0[0].Unrestricted_value,1)
	
	fmt.Println(res1)
	fmt.Println()
	fmt.Println(res1[0].Data)
}

