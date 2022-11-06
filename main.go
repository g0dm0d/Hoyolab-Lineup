package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/g0dm0d/lineup/entity"
	"github.com/g0dm0d/lineup/tools"
)

func main() {
	req := tools.GetData("")
	var data []entity.RetData
	for {
		if req.Data.Npt == "" {
			break
		}
		req = tools.GetData(req.Data.Npt)
		data = append(data, req)
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data.json", file, 0644)
}
