package main

import (
	"fmt"
)

func main() {

	v := 0
	limit := 1000

	for v < limit {
		pool := make(chan struct{}, 5)
		go func() {
			defer func() {
				<-pool
			}()

			if v < limit {
				v++
			}
		}()

	}
	fmt.Println(v)

}
