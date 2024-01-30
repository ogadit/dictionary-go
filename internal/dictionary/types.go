package dictionary

type Response struct {
	Word      string     `json:"word"`
	Phonetics []Phonetic `json:"phonetics"`
	Meanings  []Meaning  `json:"meanings"`
}

type Phonetic struct {
	Text string `json:"text"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Text    string `json:"definition"`
	Example string `json:"example"`
}
