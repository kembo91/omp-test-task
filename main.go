package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kebmo91/omp-test-task/processing"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("must pass filepath as an only parameter")
	}
	filepath := os.Args[1]
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("file %v does not exist", filepath))
	}
	tp, err := processing.NewTypeProcessor(f)
	if err != nil {
		log.Fatal(err)
	}
	p1, p2, err := tp.Process()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("max price products: %v \nmax rating products: %v", p1, p2)
}
