package tools

import (
	"github.com/g0dm0d/lineup/entity"
)

func SearchID(data entity.Data, search string) bool {
	for i := range data.List {
		if data.List[i].ID == search {
			return true
		}
	}
	return false
}
