package pkg

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"go-sherlock/config"
)

func Worker(site config.WebSites, username string, work *sync.WaitGroup) {
	for _, c := range site {
		// if strings.Contains(c.URL, "instagram") {
		// 	fmt.Println(c)
		// }
		go tester(c, username, work)
	}
}

func tester(web config.WebSite, username string, work *sync.WaitGroup) {
	web.PutUserToURL(username)
	if web.URLProbe != "" {
		web.URL = web.URLProbe
	}
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(web.URL)
	if err != nil {
		work.Done()
		return
	}
	exist := false
	switch web.ErrorType {
	case "status_code":
		exist = Status(web, resp)
	case "message":
		exist = ErrorMsg(web, resp)
	}

	if exist {
		printer(web, username)
	}
	work.Done()

}

func printer(web config.WebSite, username string) {
	fmt.Println(username + " >>>>>>    " + web.URL)
}
func Status(web config.WebSite, resp *http.Response) bool {
	switch resp.StatusCode {
	case web.ErrorCode:
		return false
	case 200:
		return true
	}
	return false

}
func ErrorMsg(web config.WebSite, resp *http.Response) bool {
	body, _ := io.ReadAll(resp.Body)
	html := string(body)
	if strings.Contains(html, fmt.Sprintf("%v", web.ErrorMsg)) {
		return false
	} else {
		return true
	}

}
