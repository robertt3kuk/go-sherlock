package config

import (
	"encoding/json"
	"os"
	"strings"
)

type WebSite struct {
	ErrorMsg  string            `json:"error_message,omitempty"`
	ErrorCode int               `json:"error_code,omitempty"`
	ErrorType string            `json:"error_type"`
	Headers   map[string]string `json:"headers,omitempty"`
	URL       string            `json:"url"`
	URLMain   string            `json:"url_main"`
	URLProbe  string            `json:"url_probe"`
	Claimed   string            `json:"claimed"`
	Unclaimed string            `json:"unclaimed"`
}

type WebSites map[string]WebSite

func ParseSites(web *WebSites) error {
	data, err := os.ReadFile("./config/data.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &web); err != nil {
		return err
	}
	return nil
}
func (web *WebSite) PutUserToURL(user string) {
	s := strings.ReplaceAll(web.URL, "{}", user)
	web.URL = s
	if web.URLProbe != "" {
		web.URLProbe = strings.ReplaceAll(web.URLProbe, "{}", user)
	}
}
