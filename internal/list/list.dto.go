package list

import (
	"slices"

	dtox "a.com/test/internal/httpx/dto"
)

type CreateListDto struct {
	Title string `json:"title" validate:"required,min=3"`
}

type UpdateListDto struct {
	Title  *string     `json:"title" validate:"omitempty,min=3"`
	Status *ListStatus `json:"status"`
}

func (dto UpdateListDto) EnumValidation() *dtox.DtoValidationFailed {
	if dto.Status == nil {
		return nil
	}

	if slices.Contains([]ListStatus{ListStatusCreated, ListStatusInProgress, ListStatusDone}, *dto.Status) {
		return nil
	}

	err := dtox.DtoValidationFailed{}.EnumFailed("Status", []string{
		string(ListStatusCreated), string(ListStatusInProgress),
		string(ListStatusDone),
	},
	)

	return &err
}
