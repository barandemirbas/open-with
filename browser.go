package openwith

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// Browser function opens a new browser tab for pointing url with default browser.
func Browser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		err = fmt.Errorf("Sorry, we do not support this OS.")
	}
	if err != nil {
		log.Fatal(err)
	}
}
