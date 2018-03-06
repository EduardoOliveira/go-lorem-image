package core

import (
	"github.com/eduardooliveira/go-lorem-image/core/server"
	"github.com/eduardooliveira/go-lorem-image/core/images"
)

func Start() {
	images.Register(server.GetGroup("lorem"))
	server.Start()
}
