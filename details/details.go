package details

import (
	"reflect"
)

type Detals struct {
	Vector []bool
	Names []int
}

func FormDetailWithoutRepetition(detailArr []Detals) []Detals{
	nextArr := []Detals{}
	for i, detail := range detailArr {
		isAdd := true
		pos := -1
		for j, newDetail := range nextArr {
			if reflect.DeepEqual(newDetail.Vector, detail.Vector) {
				isAdd = false
				pos = j
				break
			}
		}
		if isAdd {
			newDetail := Detals{Vector: detail.Vector, Names: []int{i + 1}}
			nextArr = append(nextArr, newDetail)
		} else {
			nextArr[pos].Names = append(nextArr[pos].Names, i + 1)
		}
	}
	return nextArr
}