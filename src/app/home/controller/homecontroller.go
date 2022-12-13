package example

import (
	"net/http"

	helper "github.com/faisd405/go-restapi-chi/src/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := helper.M{"title": "Example", "name": "Faisal", "age": 25}

	tpl := helper.NewRenderer("src/views/*.html", true)
	tpl.Render(w, "index.html", data, r)
}
