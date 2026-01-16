package kjbc

import (
	"fmt"
	"strings"

	"github.com/jrswab/kjbc/cons"
	"github.com/jrswab/kjbc/dicts"
)

const (
	LanguageGreek  = "greek"
	LanguageHebrew = "hebrew"
)

type EntryRef struct {
	Number string   `json:"number"`
	Verses []string `json:"verses,omitempty"`
}

type WordResult struct {
	Word     string     `json:"word"`
	Language string     `json:"language"`
	Entries  []EntryRef `json:"entries"`
}

type EntryResult struct {
	Number     string `json:"number"`
	Language   string `json:"language"`
	Word       string `json:"word"`
	Pronounce  string `json:"pronounce,omitempty"`
	Definition string `json:"definition"`
}

func NormalizeLanguage(lang string) (string, error) {
	if lang == "" {
		return LanguageGreek, nil
	}

	normalized := strings.ToLower(lang)
	switch normalized {
	case LanguageGreek, LanguageHebrew:
		return normalized, nil
	default:
		return "", fmt.Errorf("unsupported language: %s", lang)
	}
}

func SearchWord(lang, word string, includeVerses bool) (WordResult, error) {
	language, err := NormalizeLanguage(lang)
	if err != nil {
		return WordResult{}, err
	}

	refs, err := cons.Find(language, word)
	if err != nil {
		return WordResult{}, err
	}

	entries := make([]EntryRef, 0, len(refs))
	for _, ref := range refs {
		entry := EntryRef{Number: ref.Number}
		if includeVerses {
			entry.Verses = ref.Verses
		}
		entries = append(entries, entry)
	}

	return WordResult{
		Word:     word,
		Language: language,
		Entries:  entries,
	}, nil
}

func GetEntry(lang, number string) (EntryResult, error) {
	language, err := NormalizeLanguage(lang)
	if err != nil {
		return EntryResult{}, err
	}

	entry, err := dicts.Get(language, number)
	if err != nil {
		return EntryResult{}, err
	}

	return EntryResult{
		Number:     number,
		Language:   language,
		Word:       entry.Word,
		Pronounce:  entry.Pronounce,
		Definition: entry.Definition,
	}, nil
}

func Languages() []string {
	return []string{LanguageGreek, LanguageHebrew}
}
