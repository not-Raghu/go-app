package helpers

import (
	"time"

	"github.com/goombaio/namegenerator"
)

func GeneateNames() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()
	return name
}
