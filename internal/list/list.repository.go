package list

import (
	"context"

	"a.com/test/internal/database"
)

func ReadAllListRepository(ctx context.Context) ([]List, error) {
	repo := database.GetRepo()
	items := []List{}

	if err := repo.WithContext(ctx).Order("id").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func SaveListRepository(ctx context.Context, l List) (List, error) {
	repo := database.GetRepo()

	if err := repo.WithContext(ctx).Create(&l).Error; err != nil {
		return List{}, err
	}

	return l, nil
}

func UpdateListRepository(ctx context.Context, id int, l List) (List, error) {
	repo := database.GetRepo()

	r := repo.WithContext(ctx).Where("id = ?", id).Updates(&l).First(&l)

	if r.Error != nil {
		return List{}, r.Error
	}

	return l, nil
}

func DeleteListRepository(ctx context.Context, id int) error {
	repo := database.GetRepo()

	if err := repo.WithContext(ctx).Where("id = ?", id).Delete(&List{}).Error; err != nil {
		return err
	}

	return nil
}
