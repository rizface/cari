package film

import (
	"cari/cmd/helper"
)

const (
	REBAHIN = "http://167.88.14.147/?s="
	REBAHINBYDATE = "http://167.88.14.147/release-year/"
	SELECTOR = ".ml-item"
)


func SearchFilm(title string) {
	body := helper.Request(REBAHIN+title)
	defer body.Close()
	films := helper.ExtractContent(SELECTOR,body)
	helper.Listing(title,films)
}

func SearchByYear(year string) {
	body := helper.Request(REBAHINBYDATE+year)
	defer body.Close()
	films := helper.ExtractContent(SELECTOR,body)
	helper.Listing(year + "Releases",films)
}