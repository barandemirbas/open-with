package openwith

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
)

// Browser function opens a new browser tab for pointing url with default browser.
func Browser(url string, port ...int) {
	var err error
	if port != nil {
		url += ":" + strconv.Itoa(port[0])
	}
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
