package main

import (
	"fmt"

	"github.com/mbraeutig/zeit.co/api"
)

func main() {
	fmt.Println("main()")
	api.Mon = "main()"
}

func init() {
	api.Str = "haHa"
}
