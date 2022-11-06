package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/g0dm0d/lineup/entity"
)

func GetData(npt string) (data entity.RetData) {
	link := "https://sg-public-api.hoyoverse.com/event/simulatoros/lineup/index?"
	params := fmt.Sprintf("next_page_token=%s&order=CreatedTime&tag_id=2", url.QueryEscape(npt))
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, link+params, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data = ParseData(body)
	return
}
