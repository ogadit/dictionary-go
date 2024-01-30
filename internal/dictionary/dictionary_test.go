package dictionary_test

import (
	"testing"

	"osamagadit.com/dict/internal/dictionary"
)

func TestGetPhonetic(t *testing.T) {
	resp := dictionary.Response{
		Word:      "hello",
		Phonetics: []dictionary.Phonetic{},
	}

	_, err := resp.GetPhonetic()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	resp2 := dictionary.Response{
		Word: "hello",
		Phonetics: []dictionary.Phonetic{
			{Text: "/həˈləʊ/"},
			{Text: "/həˈloʊ/"},
		},
	}
	phonetic2, err2 := resp2.GetPhonetic()
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if phonetic2 != "/həˈləʊ/" {
		t.Errorf("Expected /həˈləʊ/, got %v", phonetic2)
	}

	resp3 := dictionary.Response{
		Word: "hello",
		Phonetics: []dictionary.Phonetic{
			{},
			{Text: "/həˈləʊ/"},
			{Text: "/həˈloʊ/"},
		},
	}

	phonetic3, err3 := resp3.GetPhonetic()
	if err3 != nil {
		t.Errorf("Unexpected error %v", err3)
	}

	if phonetic3 != "/həˈləʊ/" {
		t.Errorf("Expected /həˈləʊ/, got %v", phonetic3)
	}
}
