package main

import "testing"

func TestOk(t *testing.T) {
	t.Log("this test passes")
}

func TestFail(t *testing.T) {
	t.Error("this test fails")
}
