package dtox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var v = validator.New(validator.WithRequiredStructEnabled())

type DtoValidationFailed struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func PathParser(path string) (int, DtoValidationFailed) {
	data, parseError := strconv.Atoi(path)

	if parseError != nil {
		return 0, DtoValidationFailed{
			Field: path,
			Tag:   "wrong path",
		}
	}

	return data, DtoValidationFailed{}
}

func BodyParser[D any](reqData io.Reader, data D) (D, DtoValidationFailed) {
	parseError := json.NewDecoder(reqData).Decode(&data)

	var zero D

	if parseError != nil {
		return zero, DtoValidationFailed{
			Field: parseError.Error(),
			Tag:   parseError.Error(),
		}
	}

	return data, DtoValidationFailed{}
}

func ValidateDto[D any](
	reqData io.Reader, dto D, parser func(io.Reader, D) (D, DtoValidationFailed),
) (D, []DtoValidationFailed, error) {
	data, parseError := parser(reqData, dto)

	var zero D
	if parseError.Field != "" {
		return zero, []DtoValidationFailed{parseError}, nil
	}

	if err := v.Struct(data); err != nil {
		var validationErrors validator.ValidationErrors

		if !errors.As(err, &validationErrors) {
			return zero, nil, fmt.Errorf("validation server error: %w", err)
		}

		var failedValidations []DtoValidationFailed

		for _, fe := range validationErrors {
			failedValidations = append(failedValidations, DtoValidationFailed{
				Field: fe.Field(),
				Tag:   fe.Tag(),
			})
		}

		return zero, failedValidations, nil
	}
	if val, ok := any(data).(WithEnumValidation); ok {
		if enumErr := val.EnumValidation(); enumErr != nil {
			return zero, []DtoValidationFailed{*enumErr}, nil
		}
	}

	return data, nil, nil
}

type WithEnumValidation interface {
	EnumValidation() *DtoValidationFailed
}

func (dvf DtoValidationFailed) EnumFailed(field string, valids []string) DtoValidationFailed {
	tag := field + " should be: "

	for _, valid := range valids {
		tag += valid + ", "
	}

	return DtoValidationFailed{
		Tag:   tag,
		Field: field,
	}
}
