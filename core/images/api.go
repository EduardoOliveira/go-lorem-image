package images

import (
	"github.com/labstack/echo"
	"time"
	"math/rand"
	"strconv"
	"github.com/disintegration/imaging"
	"log"
	"net/http"
	"image/color"
	"image"
)

func ServeImage(c echo.Context) error {
	dir := c.Param("dir")
	if dir == "" {
		return c.NoContent(http.StatusNotFound)
	}

	if files[dir] == nil || len(files[dir]) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	idx := r.Intn(len(files[dir]) - 1)
	log.Println(files[dir][idx])
	log.Println(idx)

	src, err := imaging.Open(files[dir][idx])
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	var h = 0
	var w = 0

	if c.Param("h") != "" {
		h, _ = strconv.Atoi(c.Param("h"))
	}
	if c.Param("w") != "" {
		w, _ = strconv.Atoi(c.Param("w"))
	}
	dst := src

	if w != 0 && h != 0 {
		src = imaging.Fill(src, w, h, imaging.Center, imaging.Lanczos)
		dst = imaging.New(w, h, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, src, image.Pt(0, 0))
	} else if w != h && (w == 0 || h == 0) {
		dst = imaging.Resize(src, w, h, imaging.Lanczos)
	}
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Response().Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Response().Header().Set("Expires", "0")                                         // Proxies.
	imaging.Encode(c.Response().Writer, dst, imaging.JPEG)
	c.Response().Flush()

	return nil
}
