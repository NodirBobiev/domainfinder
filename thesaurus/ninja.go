package thesaurus

import (
	"encoding/json"
	"net/http"
)

type ninjaResponse struct {
	Synonyms []string `json:"synonyms"`
	Antonyms []string `json:"antonyms"`
}

type Ninja struct {
	APIKey string
}

func (n *Ninja) Synonyms(term string) ([]string, error) {
	whoisServer := "https://api.api-ninjas.com/v1/thesaurus?word=" + term
	client := &http.Client{}
	req, err := http.NewRequest("GET", whoisServer, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("X-Api-Key", n.APIKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	var data ninjaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []string{}, err
	}
	return data.Synonyms, nil
}
