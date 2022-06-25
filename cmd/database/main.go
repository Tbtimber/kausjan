package main

import (
	"log"

	"github.com/Tbtimber/kausjan/model"
)

func main() {
	neo4jDriver, err := model.ParseConfiguration().NewDriver()
	defer model.UnsafeClose(neo4jDriver)
	if err != nil {
		log.Fatal(err)
	}

}
