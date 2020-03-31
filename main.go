package main

import (
	"github.com/mbraeutig/zeit.co/api"
)

func main() {
	api.RootMain = "calling from root main()"
}

func init() {
	api.RootInit = "calling from root init()"
}
