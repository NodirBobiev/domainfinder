package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

type BigHuge struct {
	APIKey string
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: Failed when looking for synonyms for \"%s\" %s", term, err)

	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}

// b9d26460663d4a330c0b1f9a148c9e66
