package main

import (
	"log"

	"github.com/zztkm/toolformation"
)

// TODO: いずれ cobra かなんか使ったCLIにする1
func main() {
	t, err := toolformation.New("ToolFormation.yml")
	if err != nil {
		log.Fatal(err)
	}
	t.Install()
}
