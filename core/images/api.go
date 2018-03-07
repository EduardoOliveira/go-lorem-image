package images

import (
	"github.com/labstack/echo"
	"time"
	"math/rand"
	"strconv"
	"github.com/disintegration/imaging"
	"log"
)

func ServeImage(c echo.Context) error {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	idx := r.Intn(len(files) - 1)
	log.Println(files[idx])
	log.Println(idx)

	src, err := imaging.Open(files[idx])
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}

	var h = 0
	var w = 0

	if c.Param("h") != "" {
		h, _ = strconv.Atoi(c.Param("h"))
	}
	if c.Param("w") != "" {
		w, _ = strconv.Atoi(c.Param("w"))
	}
	if w != 0 && h != 0 {
		src = imaging.CropAnchor(src, w, h, imaging.Center)
	} else if w != h && (w == 0 || h == 0) {
		src = imaging.Resize(src, w, h, imaging.Lanczos)
	}
	c.Response().Header().Set(echo.HeaderContentType, "Application/Jpeg")
	imaging.Encode(c.Response().Writer, src, imaging.JPEG)
	c.Response().Flush()

	return nil
}
