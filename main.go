package main

import (
	"fmt"
	"gips/records"
	"os"
)

func main() {
	recs := records.FromFile(os.Args[1])
	for _, val := range recs {
		if val.IsRLE {
			fmt.Println(val)
			fmt.Println(val.Value())
			break
		}
	}
}
