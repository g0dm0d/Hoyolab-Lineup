package tools

import (
	"github.com/g0dm0d/lineup/entity"
)

func SearchID(data entity.Data, search string) bool {
	for i := 0; i < len(data.List); i++ {
		if data.List[i].ID == search {
			return true
		}
	}
	return false
}
