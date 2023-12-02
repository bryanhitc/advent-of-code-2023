package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Finished running in " + elapsed.String() + ".")
}
