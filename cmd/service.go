package cmd

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/icza/gox/stringsx"
	"gopkg.in/yaml.v2"
)

var latestContent string

type Entry struct {
	Value string
	Count int
	Date  time.Time
}

func (e Entry) DateFormatted() string {
	return e.Date.Format("02.01.2006 15:04:05")
}

func (e Entry) ShortValue() string {
	maxLen := 30
	if len(e.Value) > maxLen {
		return string(e.Value[0:maxLen])
	}
	return e.Value
}

func NewEntry(value string) Entry {
	return Entry{
		Value: value,
		Count: 0,
		Date:  time.Now(),
	}
}

type Entries struct {
	Values []Entry
}

func (entries Entries) Add(value string) *Entries {
	e := NewEntry(value)
	entries.Values = append(entries.Values, e)
	return &Entries{Values: entries.Values}
}

func List() (*Entries, error) {
	entries, err := loadEntries(plippyFile)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func Search(query string, closestN int) (*Entries, error) {

	entries, err := loadEntries(plippyFile)
	if err != nil {
		return nil, err
	}
	results := entries.Search(query, closestN)
	return results, nil

}

func WriteToClipboard(text string) error {
	return clipboard.WriteAll(text)
}

func StartPlippy() error {

	for {

		content, err := clipboard.ReadAll()
		content = stringsx.Clean(content)

		if validate(content) == false {
			continue
		}

		if latestContent == content {
			continue
		} else {
			latestContent = content
		}

		entries, err := loadEntries(plippyFile)
		entries = entries.Add(content)

		if err != nil {
			return err
		}

		err = saveContent(plippyFile, *entries)
		if err != nil {
			return err
		}
		time.Sleep(500)
	}
}

func validate(text string) bool {

	if len(text) < 5 {
		return false
	}

	return true
}

func saveContent(file string, entries Entries) error {
	y, err := yaml.Marshal(entries)
	if err != nil {
		return err
	}
	return writeRaw(file, y)
}

func loadEntries(file string) (*Entries, error) {
	fileContent, err := loadRaw(file)

	if err != nil {
		return nil, err
	}

	entries := Entries{}
	err = yaml.Unmarshal(fileContent, &entries)

	if err != nil {
		return nil, err
	}
	return &entries, nil
}

func writeRaw(file string, content []byte) error {
	err := ioutil.WriteFile(file, content, 744)
	return err
}

func loadRaw(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (entries Entries) Search(query string, closesttN int) *Entries {
	resultEntries := Entries{}
	for _, v := range entries.Values {
		if strings.Contains(v.Value, query) {
			resultEntries.Add(v.Value)
		}
	}
	return &resultEntries
}
