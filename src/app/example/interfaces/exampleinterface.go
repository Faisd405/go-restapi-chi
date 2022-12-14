package interfaces

import (
	"context"

	exampleModel "github.com/faisd405/go-restapi-chi/src/app/example/model"
)

type ExampleRepository interface {
	// function(parameterType) (returnType)
	FindAll(ctx context.Context) ([]exampleModel.Example, error)
	FindById(ctx context.Context, id uint) (exampleModel.Example, error)
	Create(ctx context.Context, example *exampleModel.Example) error
	Update(ctx context.Context, id uint, updateExample *exampleModel.Example) error
	Delete(ctx context.Context, id uint) error
}

type ExampleService interface {
	FindAll(ctx context.Context) ([]exampleModel.Example, error)
	FindById(ctx context.Context, id uint) (exampleModel.Example, error)
	Create(ctx context.Context, example *exampleModel.Example) error
	Update(ctx context.Context, id uint, updateExample *exampleModel.Example) error
	Delete(ctx context.Context, id uint) error
}
