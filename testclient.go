package main

import (
	"github.com/jefferai/gomodtest/api"
	"github.com/kr/pretty"
)

func main() {
	c, _ := api.NewClient(nil)
	pretty.Print(c)
}
