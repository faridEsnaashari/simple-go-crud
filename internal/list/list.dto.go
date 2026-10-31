package list

import (
	"fmt"
	"slices"

	dtox "a.com/test/internal/httpx/dto"
)

type CreateListDto struct {
	Title string `json:"title" validate:"required,min=3"`
}

type UpdateListDto struct {
	ID     int        `json:"id"`
	Title  string     `json:"title" validate:"min=3"`
	Status ListStatus `json:"status" validate:"min=3"`
}

func (dto UpdateListDto) EnumValidation() *dtox.DtoValidationFailed {

	fmt.Println()
	if slices.Contains([]ListStatus{ListStatusCreated, ListStatusInProgress, ListStatusDone}, dto.Status) {
		return nil
	}

	err := dtox.DtoValidationFailed{}.EnumFailed("Status", []string{
		string(ListStatusCreated), string(ListStatusInProgress),
		string(ListStatusDone),
	},
	)

	return &err
}
