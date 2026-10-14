package req

import (
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSONRes(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.JSONRes(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil

}
