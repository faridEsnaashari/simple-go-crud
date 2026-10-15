package list

import (
	"strconv"

	dtox "a.com/test/internal/httpx/dto"
	errorx "a.com/test/internal/httpx/errors"
	"a.com/test/internal/httpx/middlewares"
)

func ReadAllListHandler(req middlewares.Request) ([]List, errorx.HandledError) {
	return ReadAllListService(), errorx.HandledError{}
}

func SaveListHandler(req middlewares.Request) (List, errorx.HandledError) {
	var l CreateListDto

	data, err := dtox.ValidateDto(req.Body, l, dtox.BodyParser)

	if err != nil {
		return List{}, errorx.ValidationError(err)
	}

	return SaveListService(data), errorx.HandledError{}
}

func UpdateListHandler(req middlewares.Request) (List, errorx.HandledError) {
	var l UpdateListDto

	id, pathError := dtox.PathParser(req.PathValue("id"))

	if pathError.Field != "" {
		return List{}, errorx.ValidationError([]dtox.DtoValidationFailed{pathError})
	}

	data, err := dtox.ValidateDto(req.Body, l, dtox.BodyParser)

	if err != nil {
		return List{}, errorx.ValidationError(err)
	}

	r, e := UpdateListService(id, data)

	return r, e

}

func DeleteListHandler(req middlewares.Request) (any, errorx.HandledError) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	e := DeleteListService(id)
	return nil, e

}
