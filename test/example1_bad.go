package main

import (
	"github.com/miguelmota/gopractices/bad/example1/foo"
	"github.com/miguelmota/gopractices/bad/example1/thirdparty"
)

func main() {
	fooSvc := foo.New(foo.Props{
		MyPrivateKey: "my private key",
		Name:         "Bob",
	})

	thirdparty.NewPkg(fooSvc)
	// OUTPUT:
	// {my private key Bob}

	fooSvc.UsePrivateKey()
}
