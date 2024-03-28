package whois

import (
	"io"
	"net/http"
)

type Ninja struct {
	APIKey string
}

func (n *Ninja) Exists(domain string) (bool, error) {
	whoisServer := "https://api.api-ninjas.com/v1/whois?domain=" + domain
	client := &http.Client{}
	req, err := http.NewRequest("GET", whoisServer, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("X-Api-Key", n.APIKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if string(body) == "{}" {
		return false, nil
	}
	return true, nil
}
