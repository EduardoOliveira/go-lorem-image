package images

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/eduardooliveira/go-lorem-image/core/config"
	"github.com/labstack/echo"
)

var (
	files   []string
	imgRoot string
)

func init() {
	imgRoot = config.Config().GetString("images.root")
	dFiles, err := ioutil.ReadDir(imgRoot)
	if err != nil {
		log.Fatal(err)
	}

	files = make([]string, 1)

	for _, file := range dFiles {
		files = append(files, fmt.Sprintf("%s/%s", imgRoot, file.Name()))
	}

	log.Println(files)
}
func Register(group *echo.Group) {
	group.GET("/img", ServeImage)
	group.GET("/img/h/:h", ServeImage)
	group.GET("/img/w/:w", ServeImage)
	group.GET("/img/:h/:w", ServeImage)
}
