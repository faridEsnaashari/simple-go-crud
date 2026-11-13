package list

import (
	"context"

	"a.com/test/internal/database"

	"gorm.io/gorm"
)

func ReadAllListRepository(ctx context.Context, tx *gorm.DB) ([]List, error) {
	repo := database.GetRepo[List](tx)
	items, err := repo.Order("id").Find(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func SaveListRepository(ctx context.Context, l List, tx *gorm.DB) (List, error) {
	repo := database.GetRepo[List](tx)

	if err := repo.Create(ctx, &l); err != nil {
		return List{}, err
	}

	return l, nil
}

func UpdateListRepository(ctx context.Context, id int, l UpdateListDto, tx *gorm.DB) (List, error) {
	repo := database.GetRepo[List](tx).Table("lists")
	updateRepo := database.GetRepo[any](tx).Table("lists")

	_, e := updateRepo.Where("id=?", id).Updates(ctx, l)
	if e != nil {
		return List{}, e
	}

	list, err := repo.Where("id=?", id).First(ctx)
	if err == gorm.ErrRecordNotFound {
		return List{}, nil
	}

	if err != nil {
		return List{}, err
	}

	return list, nil
}

func DeleteListRepository(ctx context.Context, id int, tx *gorm.DB) error {
	repo := database.GetRepo[List](tx)

	if _, err := repo.Where("id = ?", id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
