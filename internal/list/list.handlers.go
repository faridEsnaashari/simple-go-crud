package list

import (
	"net/http"

	dtox "a.com/test/internal/httpx/dto"
	errorx "a.com/test/internal/httpx/errors"
)

func ReadAllListHandler(req *http.Request) ([]List, errorx.HandledError) {
	return ReadAllListService(req.Context())
}

func SaveListHandler(req *http.Request) (List, errorx.HandledError) {
	var l CreateListDto

	data, err, serverErr := dtox.ValidateDto(req.Body, l, dtox.BodyParser)
	if serverErr != nil {
		return List{}, errorx.InternalError(serverErr)
	}

	if err != nil {
		return List{}, errorx.ValidationError(err)
	}

	return SaveListService(req.Context(), data)
}

func UpdateListHandler(req *http.Request) (List, errorx.HandledError) {
	var l UpdateListDto

	id, pathError := dtox.PathParser(req.PathValue("id"))

	if pathError.Field != "" {
		return List{}, errorx.ValidationError([]dtox.DtoValidationFailed{pathError})
	}

	data, err, serverErr := dtox.ValidateDto(req.Body, l, dtox.BodyParser)

	if serverErr != nil {
		return List{}, errorx.InternalError(serverErr)
	}

	if err != nil {
		return List{}, errorx.ValidationError(err)
	}

	r, e := UpdateListService(req.Context(), id, data)

	return r, e
}

func DeleteListHandler(req *http.Request) (any, errorx.HandledError) {
	id, pathError := dtox.PathParser(req.PathValue("id"))

	if pathError.Field != "" {
		return List{}, errorx.ValidationError([]dtox.DtoValidationFailed{pathError})
	}

	e := DeleteListService(req.Context(), id)
	return nil, e
}
