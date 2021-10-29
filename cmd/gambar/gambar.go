package gambar

import (
	"cari/cmd/helper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	WALPAPERCRAFT_QUERY = "https://wallpaperscraft.com/search/?query="
	DOWNLOAD = "https://wallpaperscraft.com/download/"
)

var cat string

func VisitWeb(thread int, page int,dir string,category string) {
	start := time.Now()
	runtime.GOMAXPROCS(thread)
	cat = category
	var wg sync.WaitGroup
	for i := 1; i <= page ; i++ {
		wg = sync.WaitGroup{}
		body := helper.Request(fmt.Sprintf("https://wallpaperscraft.com/search/?order=&page=%d&query=%s&size=",i,cat))
		defer body.Close()

		doc,_ := goquery.NewDocumentFromReader(body)
		finder := doc.Find(".wallpapers__link")

		wg.Add(finder.Length())
		finder.Each(func(i int, selection *goquery.Selection) {
			link,_ := selection.Attr("href")
			go Download(dir,link, &wg)
		})
		wg.Wait()
	}
	fmt.Println(time.Since(start))
}

func Download(dir string,link string, wg *sync.WaitGroup) {
	defer wg.Done()
	items := strings.Split(link, "/")
	path := items[len(items) - 1]
	link = "https://images.wallpaperscraft.com/image/single/"+path+"_1280x720.jpg"
	body := helper.Request(link)
	defer body.Close()

	_,err := os.Stat(dir)

	if err != nil {
		helper.PrintIfError(err)
	}

	fileName := dir+ cat + "-"+strconv.Itoa(rand.Intn(int(time.Now().Unix()))) + ".jpg"
	file,_ := os.Create(fileName)
	io.Copy(file,body)
	log.Print("Selesai")
}

