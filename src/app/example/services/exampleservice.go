package services

import (
	"context"

	exampleInterface "github.com/faisd405/go-restapi-chi/src/app/example/interfaces"
	exampleModel "github.com/faisd405/go-restapi-chi/src/app/example/model"
)

type exampleService struct {
	exampleRepository exampleInterface.ExampleRepository
}

func (service *exampleService) FindAll(ctx context.Context) ([]exampleModel.Example, error) {
	return service.exampleRepository.FindAll(ctx)
}

func (service *exampleService) FindById(ctx context.Context, id uint) (exampleModel.Example, error) {
	return service.exampleRepository.FindById(ctx, id)
}

func (service *exampleService) Create(ctx context.Context, example *exampleModel.Example) error {
	return service.exampleRepository.Create(ctx, example)
}

func (service *exampleService) Update(ctx context.Context, id uint, updateExample *exampleModel.Example) error {
	return service.exampleRepository.Update(ctx, id, updateExample)
}

func (service *exampleService) Delete(ctx context.Context, id uint) error {
	return service.exampleRepository.Delete(ctx, id)
}

func NewExampleService(exampleRepository exampleInterface.ExampleRepository) exampleInterface.ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}
