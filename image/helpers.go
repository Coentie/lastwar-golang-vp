package image

import (
	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"os"
)

func CompareHashes(imgPath1, imgPath2 string) (bool, error) {
	file1, err := os.Open(imgPath1)
	if err != nil {
		return false, err
	}
	defer file1.Close()

	file2, err := os.Open(imgPath2)
	if err != nil {
		return false, err
	}
	defer file2.Close()

	img1, err := png.Decode(file1)
	if err != nil {
		return false, err
	}

	img2, err := png.Decode(file2)
	if err != nil {
		return false, err
	}

	hash1, err := goimagehash.AverageHash(img1)
	if err != nil {
		return false, err
	}

	hash2, err := goimagehash.AverageHash(img2)
	if err != nil {
		return false, err
	}

	distance, err := hash1.Distance(hash2)
	if err != nil {
		return false, err
	}

	return distance < 5, nil
}

func GetCrop(x, y, width, height int, screen image.Image) *image.NRGBA {
	cropRect := image.Rect(x, y, x+width, y+height)
	return imaging.Crop(screen, cropRect)
}

func OpenImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func Compare(templateHash *goimagehash.ImageHash, croppedHash *goimagehash.ImageHash, tolerance int) bool {
	distance, err := templateHash.Distance(croppedHash)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Hamming distance: %d\n", distance)

	return distance <= tolerance
}
