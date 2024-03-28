package thesaurus

type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}

// ChatSynonyms returns synonyms for the word "chat". This struct is only for
// testing. Because Big Huge Labs api, sometimes, is not avaible.
type ChatSynonyms struct {
}

func (s *ChatSynonyms) Synonyms(term string) ([]string, error) {
	return []string{
		"chat",
		"confab",
		"confabulation",
		"schmooze",
		"New World chat",
		"Old World chat",
		"conversation",
		"thrush",
		"wood warbler",
		"chew the fat",
		"shoot the breeze",
		"chitchat",
		"chatter",
	}, nil
}
