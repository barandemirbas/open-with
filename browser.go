package openwith

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
)

// Browser function opens a new browser tab for pointing url with default browser.
func Browser(URL string, PORT ...int) (*exec.Cmd, error) {
	var err error
	var cmd *exec.Cmd
	u, _ := url.Parse(URL)
	if u.Scheme == "" {
		u, _ = url.Parse("http://" + URL)
	}
	if PORT != nil {
		URL = u.Scheme + "://" + u.Host + ":" + strconv.Itoa(PORT[0]) + u.Path
	} else {
		URL = u.Scheme + "://" + u.Host + u.Path
	}
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", URL)
	case "darwin":
		cmd = exec.Command("open", URL)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", URL)
	default:
		err = fmt.Errorf("Sorry your OS not supports open the browser with a URL")
		if err != nil {
			log.Fatal(err)
		}
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
		return cmd, err
	}
	return cmd, nil
}
