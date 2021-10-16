package main

import (
	"fmt"
	"gips/records"
	"os"
)

func main() {
	recs := records.FromFile(os.Args[1])
	fmt.Println(recs)
	fmt.Println(recs[0].Value())
}
