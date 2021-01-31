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
	var scheme = []string{"http://", "https://"}
	var isURL = false
	for _, scheme := range scheme {
		if strings.Contains(URL, scheme) {
			isURL = true
			break
		}
	}
	if !isURL {
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
		err = fmt.Errorf("Sorry your OS not supports open the browser with a URL")
	}
	if err != nil {
		log.Fatal(err)
	}
}
