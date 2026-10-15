package list

import (
	"slices"

	errorx "a.com/test/internal/httpx/errors"
)

var lists []List

func getLists() []List {
	if len(lists) == 0 {
		lists = []List{
			{ID: 1, Title: "t1", Status: ListStatusCreated},
			{ID: 2, Title: "t2", Status: ListStatusCreated},
			{ID: 3, Title: "t3", Status: ListStatusCreated},
		}
	}

	return lists
}

func ReadAllListRepository() []List {
	return getLists()
}

func SaveListRepository(l List) List {
	lists = append(lists, l)
	return l
}

func UpdateListRepository(id int, l List) (List, errorx.HandledError) {
	index := slices.IndexFunc(lists, func(l List) bool {
		return l.ID == id
	})

	if index == -1 {
		return List{}, errorx.NotFoundError("list")
	}

	lists[index] = l

	return l, errorx.HandledError{}
}

func DeleteListRepository(id int) errorx.HandledError {
	index := slices.IndexFunc(lists, func(l List) bool {
		return l.ID == id
	})

	if index == -1 {
		return errorx.NotFoundError("list")
	}

	lists = slices.Delete(lists, index, index+1)

	return errorx.HandledError{}
}
