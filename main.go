package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/g0dm0d/lineup/entity"
	"github.com/g0dm0d/lineup/tools"
)

// World Exploration 		1
// Trounce Domains 			3
// Domain Challenges 		9
// Boss Battles 				24
// Spiral Abyss 				2
// Abyssal Moon Spire 	41

func main() {
	id := 41
	req := tools.GetData("", id)
	var data []entity.RetData
	for {
		if tools.SearchID(req.Data, "6365b24c8436de1e619428d0") {
			break
		}
		req = tools.GetData(req.Data.Npt, id)
		data = append(data, req)
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data.json", file, 0644)
}
