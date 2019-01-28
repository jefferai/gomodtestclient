package main

import (
	"github.com/jefferai/gomodtest/api"
)

func main() {
	c, _ := api.NewClient(nil)
	c.SetNamespace("foo")
}
