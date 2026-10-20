package list

import (
	errorx "a.com/test/internal/httpx/errors"
)

func ReadAllListService() []List {
	return ReadAllListRepository()
}

func SaveListService(dto CreateListDto) List {
	return SaveListRepository(List{
		Title:  dto.Title,
		Status: ListStatusCreated,
	})
}

func UpdateListService(id int, dto UpdateListDto) (List, errorx.HandledError) {
	list, err := UpdateListRepository(id, List{
		Title:  dto.Title,
		Status: dto.Status,
	})

	return list, err
}

func DeleteListService(id int) errorx.HandledError {
	return DeleteListRepository(id)
}
