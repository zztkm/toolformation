package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zztkm/toolformation"
)

const (
	name    = "taskformation"
	version = "0.1.0"
)

var (
	versionFlag = flag.Bool("version", false, "print version")
)

// TODO: いずれ cobra かなんか使ったCLIにする1
func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s %s", name, version)
		return
	}

	t, err := toolformation.New("ToolFormation.yml")
	if err != nil {
		log.Fatal(err)
	}
	t.Install()
}
