package list

import (
	"net/http"
	"strconv"

	dtox "a.com/test/internal/httpx/dto"
	errorx "a.com/test/internal/httpx/errors"
)

func ReadAllListHandler(req *http.Request) ([]List, errorx.HandledError) {
	return ReadAllListService(req.Context())
}

func SaveListHandler(req *http.Request) (List, errorx.HandledError) {
	var l CreateListDto

	data, err := dtox.ValidateDto(req.Body, l, dtox.BodyParser)
	if err != nil {
		return List{}, errorx.ValidationError(err)
	}

	return SaveListService(req.Context(), data)
}

func UpdateListHandler(req *http.Request) (int, errorx.HandledError) {
	var l UpdateListDto

	id, pathError := dtox.PathParser(req.PathValue("id"))

	if pathError.Field != "" {
		return 0, errorx.ValidationError([]dtox.DtoValidationFailed{pathError})
	}

	data, err := dtox.ValidateDto(req.Body, l, dtox.BodyParser)
	if err != nil {
		return 0, errorx.ValidationError(err)
	}

	r, e := UpdateListService(req.Context(), id, data)

	return r, e
}

func DeleteListHandler(req *http.Request) (any, errorx.HandledError) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	e := DeleteListService(req.Context(), id)
	return nil, e
}
