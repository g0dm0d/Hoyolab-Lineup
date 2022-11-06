package tools

import (
	"encoding/json"

	"github.com/g0dm0d/lineup/entity"
)

func ParseData(byteJson []byte) (data entity.RetData) {
	json.Unmarshal(byteJson, &data)
	return data
}
