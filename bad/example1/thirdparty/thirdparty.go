package thirdparty

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/miguelmota/gopractices/bad/example1/foo"
)

// NewPkg ...
func NewPkg(svc *foo.Service) {

	// Third party package can easily see your private key because it returns all the fields
	fmt.Println(svc.Props())
	// OUTPUT:
	// {my private key Bob}

	spew.Dump(svc)
	// OUTPUT:
	/*
		(*foo.Service)(0xc00000a180)({
		 props: (foo.Props) {
			MyPrivateKey: (string) (len=14) "my private key",
			Name: (string) (len=3) "Bob"
		 }
		})
	*/
}
