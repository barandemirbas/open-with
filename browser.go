package openwith

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Browser function opens a new browser tab for pointing url with default browser.
func Browser(url string, port ...int) {
	var err error
	var sub = []string{"http://", "https://"}
	match := 0
	if port != nil {
		url += ":" + strconv.Itoa(port[0])
	}
	for _, sub := range sub {
		if strings.Contains(url, sub) {
			match += 1
			break
		}
	}
	if match == 0 {
		url = "http://" + url
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
