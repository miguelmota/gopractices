package foo

import (
	"fmt"
	"time"
)

// example temp holder of private data
var private = make(map[int64]string)

// Foo ...
type Foo struct {
	id   int64
	name string
}

// Config ...
type Config struct {
	MyPrivateKey string
	Name         string
}

// NewFoo ...
func NewFoo(config *Config) *Foo {
	id := time.Now().UnixNano()
	private[id] = config.MyPrivateKey

	return &Foo{
		id:   id,
		name: config.Name,
	}
}

// Name ...
func (f *Foo) Name() string {
	return f.name
}

// UsePrivateKey ...
func (f *Foo) UsePrivateKey() {
	fmt.Println(len(private[f.id]))
}
