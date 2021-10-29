package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	start := time.Now()
	bucket := make([]byte, 500)
	mutext := sync.Mutex{}
	//fariz := "fa;dalkshfkladhfskjlfhgksjdhfkshdf lsd fslkdfj skldfj skldfj slkdfj skdlf jsdklf jsdkflsj dfkl sjd"
	resp,_ := http.Get("https://images.unsplash.com/photo-1624486278827-c32dbed03803?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&dl=ali-karimiboroujeni-DDCRb4i8lZg-unsplash.jpg")
	//fmt.Println(resp.Body)
	//tes,_ := io.ReadFull(strings.NewReader(fariz), bucket)
	//fmt.Println(string(bucket),tes)
	file,_ := os.Create("fariz.jpg")
	for  {
		n,_ := io.ReadFull(resp.Body,bucket)
		if n == 0 {
			break
		} else {
			mutext.Lock()
			go file.Write(bucket)
			mutext.Unlock()
			log.Println("Berhasil", n	)
		}
	}
	fmt.Println(time.Since(start))
}


func TestDownloadUnsplash(t *testing.T) {
	url := "https://wallpaperscraft.com/search/?query=nature"
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		body := resp.Body
		defer body.Close()
		doc,_:=goquery.NewDocumentFromReader(body)
		doc.Find(".wallpapers__link").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Attr("href"))
		})
	}


	//url := "https://wallpaperscraft.com/wallpaper/balls_flying_sky_nature_91285"
	//resp,err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	body := resp.Body
	//	defer body.Close()
	//	doc,err := goquery.NewDocumentFromReader(body)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		doc.Find(".resolutions__link").Each(func(i int, selection *goquery.Selection) {
	//			if selection.Text() == "1280x720" {
	//				link,_ := selection.Attr("href")
	//				fmt.Println(link)
	//				resp2,_ := http.Get("https://wallpaperscraft.com"+link)
	//				defer resp2.Body.Close()
	//				doc2,_ := goquery.NewDocumentFromReader(resp2.Body)
	//				doc2.Find(".gui-button").Each(func(i int, selection *goquery.Selection) {
	//					fmt.Println(selection.Attr("href"))
	//				})
	//			}
	//		})
	//	}
	//}

}

