package ffmpeg

import (
	_ "embed"
	"io/ioutil"
	"time"
)

//go:embed ffmpeg
var b []byte

var loaded bool

func Binary() string {
	if !loaded {
		loaded = true
		ioutil.WriteFile("./ffmpeg", b, 0777)
	}
	return "./ffmpeg"
}

func Clean() {
	os.Remove("./ffmpeg")
}