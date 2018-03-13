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
	files, err := load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(files)
}

func Register(group *echo.Group) {
	group.PATCH("/reload", Reload)
	group.GET("/:dir", ServeImage)
	group.GET("/:dir/h/:h", ServeImage)
	group.GET("/:dir/w/:w", ServeImage)
	group.GET("/:dir/:h/:w", ServeImage)
}

func load() (rtn map[string][]string, err error) {
	rtn = make(map[string][]string, 1)

	root, err := ioutil.ReadDir(imgRoot)
	if err != nil {
		return
	}

	for _, dir := range root {
		if dir.IsDir() {
			d, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", imgRoot, dir.Name()))
			if err != nil {
				break
			}
			for _, f := range d {
				if rtn[dir.Name()] == nil {
					rtn[dir.Name()] = make([]string, 1)
				}
				rtn[dir.Name()] = append(rtn[dir.Name()], fmt.Sprintf("%s/%s/%s", imgRoot, dir.Name(), f.Name()))
			}
		}
	}
	return
}
