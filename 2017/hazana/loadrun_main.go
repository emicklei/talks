package main

import (
	"log"

	"github.com/emicklei/hazana"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	hazana.PrintSummary(hazana.Run(new(clockAttack), hazana.ConfigFromFlags()))
}
