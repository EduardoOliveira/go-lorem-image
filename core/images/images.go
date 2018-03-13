package images

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/eduardooliveira/go-lorem-image/core/config"
	"github.com/labstack/echo"
)

var (
	files   map[string][]string
	imgRoot string
)

func init() {
	imgRoot = config.Config().GetString("images.root")
	root, err := ioutil.ReadDir(imgRoot)
	if err != nil {
		log.Fatal(err)
	}

	files = make(map[string][]string, 1)

	for _, dir := range root {
		if dir.IsDir() {
			d, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", imgRoot, dir.Name()))
			if err != nil {
				break
			}
			for _, f := range d {
				if files[dir.Name()] == nil {
					files[dir.Name()] = make([]string, 1)
				}
				files[dir.Name()] = append(files[dir.Name()], fmt.Sprintf("%s/%s/%s", imgRoot, dir.Name(), f.Name()))
			}
		}
	}

	log.Println(files)
}
func Register(group *echo.Group) {
	group.GET("/:dir", ServeImage)
	group.GET("/:dir/h/:h", ServeImage)
	group.GET("/:dir/w/:w", ServeImage)
	group.GET("/:dir/:h/:w", ServeImage)
}
