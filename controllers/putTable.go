package controllers

import (
	"fmt"
	"kuber/repo"
	"net/http"
)

func PutData(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, repo.PutTable())
}
