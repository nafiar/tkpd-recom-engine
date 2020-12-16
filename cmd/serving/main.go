package main

import (
	"log"

	serving "github.com/nafiar/tkpd-recom-engine/internal/serving"
)

func main() {
	log.Println(serving.Run())
}
