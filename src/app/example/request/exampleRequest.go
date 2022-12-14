package request

type ExampleRequest struct {
	Example1 string `json:"example1" validate:"required"`
	Example2 string `json:"example2" validate:"required"`
}
