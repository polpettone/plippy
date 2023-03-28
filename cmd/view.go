package cmd

import (
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func ShowSelectionView(entries Entries) (*Entry, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .DateFormatted | blue }} {{ .ShortValue | green }} {{ .Count | green }}",
		Inactive: "{{ .DateFormatted | blue }} {{ .ShortValue | black }} {{ .Count | black }}",
		Selected: "\U00001B61 {{ .DateFormatted | red }} {{ .ShortValue | red }} {{ .Count | red }}",

		Details: `
--------- Entry ----------
{{ "Value:" | faint }}	{{ .Value }}
{{ "Date:" | faint }}	{{ .DateFormatted }}
{{ "Count:" | faint }}	{{ .Count }}`,
	}

	searcher := func(input string, index int) bool {
		entry := entries.Values[index]
		name := strings.Replace(strings.ToLower(entry.Value), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Entries",
		Items:     entries.Values,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {

		if err == promptui.ErrInterrupt {
			os.Exit(1)
		}

		return nil, err
	}

	return &entries.Values[i], nil
}
