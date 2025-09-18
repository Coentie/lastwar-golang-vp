package lastwar

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"lastwar/notifier/hdadb"
	"lastwar/notifier/image"
	"os"
	"path/filepath"
	"time"
)

const CONQUERER_STATE_HEADLINE = "states/conquerer-headline.png"
const CLOSE_TEMPLATE = "states/close-template.png"

func ResetState() {
	hdadb.PrintScreen()

	// Tap to click away everything else.
	_ = hdadb.Tap("700", "500")
	time.Sleep(500 * time.Millisecond)

	_ = hdadb.SwipeBottom()
	time.Sleep(500 * time.Millisecond)

	stateImages := []string{
		"instruction_state.png",
	}

	_ = hdadb.Tap("720", "500")

	for IsStuck(hdadb.PRINTSCREEN_FILE_NAME) {
		_ = Close()
		hdadb.PrintScreen()
		time.Sleep(1 * time.Second)
	}

	// Loop through each state image
	for _, stateImage := range stateImages {
		statePath := filepath.Join(hdadb.APPLICATION_STATES_FOLDER, stateImage)

		// Compare images
		same, err := image.CompareHashes(hdadb.PRINTSCREEN_FILE_NAME, statePath)
		if err != nil {
			fmt.Printf("Error comparing with %s: %v\n", statePath, err)
			continue
		}

		if same {
			_ = Close()
			break
		} else {
		}
	}
}

func ConqueredState() bool {
	if os.Getenv("CONQUERED_STATE") == "1" {
		return true
	}

	_ = hdadb.Swipe()
	time.Sleep(1 * time.Second)
	hdadb.PrintScreen()
	screenshot, err := image.OpenImage(hdadb.PRINTSCREEN_FILE_NAME)
	if err != nil {
		panic(err)
	}

	template, err := image.OpenImage(CONQUERER_STATE_HEADLINE)
	if err != nil {
		panic(err)
	}

	cropped := image.GetCrop(262, 114, 194, 45, screenshot)

	// Generate hashes
	templateHash, err := goimagehash.PerceptionHash(template)
	if err != nil {
		panic(err)
	}

	croppedHash, err := goimagehash.PerceptionHash(cropped)
	if err != nil {
		panic(err)
	}

	return image.Compare(templateHash, croppedHash, 10)
}

func ConquererState() bool {
	if os.Getenv("CONQUERER_STATE") == "1" {
		return true
	}

	_ = hdadb.Swipe()
	time.Sleep(1 * time.Second)
	hdadb.PrintScreen()
	screenshot, err := image.OpenImage(hdadb.PRINTSCREEN_FILE_NAME)
	if err != nil {
		panic(err)
	}

	template, err := image.OpenImage(CONQUERER_STATE_HEADLINE)
	if err != nil {
		panic(err)
	}

	cropped := image.GetCrop(262, 114, 194, 45, screenshot)

	// Generate hashes
	templateHash, err := goimagehash.PerceptionHash(template)
	if err != nil {
		panic(err)
	}

	croppedHash, err := goimagehash.PerceptionHash(cropped)
	if err != nil {
		panic(err)
	}

	return image.Compare(templateHash, croppedHash, 10)
}

func IsStuck(imgFile string) bool {
	screenshot, err := image.OpenImage(imgFile)
	if err != nil {
		panic(err)
	}

	template, err := image.OpenImage(CLOSE_TEMPLATE)
	if err != nil {
		panic(err)
	}

	cropped := image.GetCrop(635, 105, 60, 50, screenshot)

	templateHash, err := goimagehash.PerceptionHash(template)
	if err != nil {
		panic(err)
	}

	croppedHash, err := goimagehash.PerceptionHash(cropped)
	if err != nil {
		panic(err)
	}

	return image.Compare(templateHash, croppedHash, 10)
}
