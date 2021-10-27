package anime

import (
	"cari/cmd/helper"
)

const (
	KUSONIME ="https://kusonime.com/?s="
	SELECTOR = ".episodeye"
)


func SearchAnime(title string) {
	body := helper.Request(KUSONIME+title)
	defer body.Close()
	anime := helper.ExtractContent(SELECTOR,body)
	helper.Listing(title + " List", anime)
}
