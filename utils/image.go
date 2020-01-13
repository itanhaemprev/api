package utils

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ImageReduce(i *os.File) error {
	// decode jpeg into image.Image
	img, err := jpeg.Decode(i)
	if err != nil {
		return err
	}
	name := i.Name()
	i.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(300, 180, img, resize.Lanczos3)

	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	if err := jpeg.Encode(out, m, nil); err != nil {
		return err
	}
	return nil
}
