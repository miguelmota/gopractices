package main

import (
	"github.com/miguelmota/gopractices/better/example1/foo"
	"github.com/miguelmota/gopractices/better/example1/thirdparty"
)

func main() {
	myFoo := foo.NewFoo(&foo.Config{
		MyPrivateKey: "my private key",
		Name:         "Bob",
	})

	thirdparty.NewPkg(myFoo)
	// OUTPUT:
	// &{1545972769151596000 Bob}

	myFoo.UsePrivateKey()
}
