package main

import (
	"log"

	"github.com/zztkm/toolformation"
)

func main() {
	t, err := toolformation.New("toolformation.yml")
	if err != nil {
		log.Fatal(err)
	}
	t.Install()
}
