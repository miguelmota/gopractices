package thirdparty

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/miguelmota/gopractices/better/example1/foo"
)

// NewPkg ...
func NewPkg(myFoo *foo.Foo) {

	// Third party package cannot easily see your private key
	fmt.Println(myFoo)
	// OUTPUT:
	// &{1545972769151596000 Bob}

	spew.Dump(myFoo)
	// OUTPUT:
	/*
		(*foo.Foo)(0xc00009e140)({
		 id: (int64) 1545972769151596000,
		 name: (string) (len=3) "Bob"
		})
	*/

	// Third party forced to call individual allowed getters
	fmt.Println(myFoo.Name())

	// fmt.Println(myFoo.MyPrivateKey()) // doesn't exist, not possible
}
