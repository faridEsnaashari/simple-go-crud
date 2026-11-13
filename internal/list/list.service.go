package list

import (
	"context"

	errorx "a.com/test/internal/httpx/errors"
)

func ReadAllListService(ctx context.Context) ([]List, errorx.HandledError) {
	r, e := ReadAllListRepository(ctx, nil)
	if e != nil {
		return nil, errorx.DatabaseError(e)
	}

	return r, errorx.HandledError{}
}

func SaveListService(ctx context.Context, dto CreateListDto) (List, errorx.HandledError) {
	r, e := SaveListRepository(ctx, List{
		Title:  dto.Title,
		Status: ListStatusCreated,
	}, nil)

	if e != nil {
		return List{}, errorx.DatabaseError(e)
	}

	return r, errorx.HandledError{}
}

func UpdateListService(ctx context.Context, id int, dto UpdateListDto) (List, errorx.HandledError) {
	r, e := UpdateListRepository(ctx, id, UpdateListDto{
		Title:  dto.Title,
		Status: dto.Status,
	}, nil)

	if e != nil {
		return List{}, errorx.DatabaseError(e)
	}

	return r, errorx.HandledError{}
}

func DeleteListService(ctx context.Context, id int) errorx.HandledError {
	e := DeleteListRepository(ctx, id, nil)
	if e != nil {
		return errorx.DatabaseError(e)
	}
	return errorx.HandledError{}
}
