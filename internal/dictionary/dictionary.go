package dictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
fetchResponse returns the first response as struct from dictionaryapi.dev
*/
func FetchResponse(word string) (Response, error) {
	var response []Response

	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)

	res, err := http.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("failed to get a response from server: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("failed to fetch response, status code %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return Response{}, fmt.Errorf("failed to parse json: %v", err)
	}

	return response[0], nil
}

// GetPhonetic returns the first phonetic text from the response.
// If the phonetics slice is empty or contains elements without a text value,
// it returns an empty string and an error indicating the issue.
func (r *Response) GetPhonetic() (string, error) {
	if len(r.Phonetics) == 0 {
		return "", fmt.Errorf("no phonetic found in the response")
	}

	for _, phonetic := range r.Phonetics {
		if phonetic.Text != "" {
			return phonetic.Text, nil
		}
	}

	return "", fmt.Errorf("no phonetic text found in the response")
}
