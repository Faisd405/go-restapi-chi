package repositories

import (
	"context"

	exampleInterface "github.com/faisd405/go-restapi-chi/src/app/example/interfaces"
	exampleModel "github.com/faisd405/go-restapi-chi/src/app/example/model"
	"github.com/faisd405/go-restapi-chi/src/helper/database"
	"gorm.io/gorm"
)

type exampleRepository struct {
	DB *gorm.DB
}

func (repository *exampleRepository) FindAll(ctx context.Context, params map[string]string) ([]exampleModel.Example, error) {
	var example = []exampleModel.Example{}

	result := repository.DB.Scopes(database.Paginate(params)).Find(&example)
	return example, result.Error
}

func (repository *exampleRepository) FindById(ctx context.Context, id uint) (exampleModel.Example, error) {
	var example exampleModel.Example
	result := repository.DB.First(&example, id)
	return example, result.Error
}

func (repository *exampleRepository) Create(ctx context.Context, example *exampleModel.Example) error {
	result := repository.DB.Create(example)
	return result.Error
}

func (repository *exampleRepository) Update(ctx context.Context, id uint, updateExample *exampleModel.Example) error {
	result := repository.DB.Model(&exampleModel.Example{}).Where("id = ?", id).Updates(updateExample)
	return result.Error
}

func (repository *exampleRepository) Delete(ctx context.Context, id uint) error {
	result := repository.DB.Delete(&exampleModel.Example{}, id)
	return result.Error
}

func NewExampleRepository(DB *gorm.DB) exampleInterface.ExampleRepository {
	return &exampleRepository{DB: DB}
}
