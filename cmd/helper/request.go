package helper

import (
	"io"
	"net/http"
	"time"
)

func Request(url string) io.ReadCloser {
	client := new(http.Client)
	client.Timeout = 7 * time.Second
	request,err := http.NewRequest(http.MethodGet,url,nil)
	PrintIfError(err)
	resp,err := client.Do(request)
	PrintIfError(err)
	body := resp.Body
	return body
}
