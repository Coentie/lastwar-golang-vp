package hdadb

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const PRINTSCREEN_FILE_NAME = "screen.png"
const APPLICATION_STATES_FOLDER = "states"

func Connect() error {
	cmd := exec.Command("HD-Adb.exe", "connect", GetConnectionString())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ADB connect failed: %v | Output: %s", err, string(output))
	}
	log.Printf("ADB connected successfully: %s", string(output))
	return nil
}

func PrintScreen() {
	cmd := exec.Command("HD-Adb.exe", "-s", GetConnectionString(), "shell", "screencap", "-p", fmt.Sprintf("sdcard/%s", PRINTSCREEN_FILE_NAME))
	_, _ = cmd.CombinedOutput()

	cmd = exec.Command("HD-Adb.exe", "-s", GetConnectionString(), "pull", fmt.Sprintf("/sdcard/%s", PRINTSCREEN_FILE_NAME))
	_, _ = cmd.CombinedOutput()
}

func Tap(x, y string) error {
	x, y = applyOffset(x, y)

	cmd := exec.Command("HD-Adb.exe", "-s", GetConnectionString(), "shell", "input", "tap", x, y)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Tap command failed at (%s,%s): %v | Output: %s", x, y, err, string(output))
	}
	return nil
}

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
}

func Swipe() error {
	return exec.Command("HD-Adb.exe", "-s", GetConnectionString(), "shell", "input", "touchscreen", "swipe",
		"500", "300",
		"500", "1300",
		"300").Run()
}

func SwipeBottom() error {
	return exec.Command("HD-Adb.exe", "-s", GetConnectionString(), "shell", "input", "touchscreen", "swipe",
		"500", "300",
		"500", "100",
		"300").Run()
}

func applyOffset(x, y string) (string, string) {
	xInt, err := strconv.Atoi(x)
	if err != nil {
		return x, y
	}

	yInt, err := strconv.Atoi(y)
	if err != nil {
		return x, y
	}

	rand.Seed(time.Now().UnixNano())
	xInt += rand.Intn(5) + 1
	yInt += rand.Intn(5) + 1

	xStr := strconv.Itoa(xInt)
	yStr := strconv.Itoa(yInt)

	return xStr, yStr
}
