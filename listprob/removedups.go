package listprob

import "go_playground/list"

type listItemType interface{}

func RemoveDups(dup *list.List) *list.List {
	itemMap := make(map[listItemType]bool)
	it := dup.Iterator()
	noDups := new(list.List)

	for it.HasNext() {
		val := it.Next()
		if _, ok := itemMap[val]; !ok {
			noDups.Append(val)
			itemMap[val] = true
		}
	}
	return noDups
}
