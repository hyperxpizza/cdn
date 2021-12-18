package tests

import (
	"flag"
	"log"
	"testing"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/stretchr/testify/assert"
)

var path = flag.String("path", "", "path to config.json file")

// go test -v
func TestLoadConfig(t *testing.T) {

	flag.Parse()

	if *path == "" {
		log.Println("Config path is empty")
		t.Fail()
		return
	}

	c, err := config.NewConfig(*path)
	assert.NoError(t, err)
	c.PrettyPrint()
}
