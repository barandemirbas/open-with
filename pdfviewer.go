package openwith

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// PDFViewer function opens a specific .pdf files.
func PDFViewer(path string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", path).Start()
	case "darwin":
		err = exec.Command("open", path).Start()
	case "windows":
		err = exec.Command("explorer.exe", path).Start()
	default:
		err = fmt.Errorf("Sorry, we do not support this OS.")
	}
	if err != nil {
		log.Fatal(err)
	}
}
