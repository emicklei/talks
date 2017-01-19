package main

import (
	"errors"

	"github.com/emicklei/tre"
)

// START OMIT
func main() {
	err := doSomething("demo")
	println(tre.New(err, "doSomething failed").Error())
}

func doSomething(with string) error {
	err := doAnotherThingThatCanFail(with)
	return tre.New(err, "doAnotherThingThatCanFail failed", "with", with) // pass error, message and context
}

func doAnotherThingThatCanFail(with string) error {
	return errors.New("something bad happened")
}

// END OMIT
