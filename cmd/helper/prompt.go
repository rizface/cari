package helper

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

func Listing(title string, contents []Content) {
	templates := promptui.SelectTemplates{
		Active:   `> {{.Title | cyan | bold}}`,
		Inactive: `{{ .Title | cyan }}`,
		Selected: `{{.Title | cyan | bold}}`,
	}
	prompt := promptui.Select{
		Label:     strings.ToUpper(title) + " List",
		Items:     contents,
		Size:      len(contents),
		Templates: &templates,
	}
	_, result, err := prompt.Run()
	PrintIfError(err)
	fmt.Println(result)
	items := strings.Split(result, " ")
	link := strings.Split(items[len(items)-1], "}")[0]
	OpenBrowser(link)
}
