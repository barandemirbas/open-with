package openwith

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Browser function opens a new browser tab for pointing url with default browser.
func Browser(URL string, PORT ...int) {
	var err error
	var sub = []string{"http://", "https://"}
	match := 0
	for _, sub := range sub {
		if strings.Contains(URL, sub) {
			match += 1
			break
		}
	}
	if match == 0 {
		URL = "http://" + URL
	}
	u, _ := url.Parse(URL)
	if PORT != nil {
		URL = u.Scheme + "://" + u.Host + ":" + strconv.Itoa(PORT[0]) + u.Path
	}
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", URL).Start()
	case "darwin":
		err = exec.Command("open", URL).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", URL).Start()
	default:
		err = fmt.Errorf("Sorry, we do not support this OS.")
	}
	if err != nil {
		log.Fatal(err)
	}
}
