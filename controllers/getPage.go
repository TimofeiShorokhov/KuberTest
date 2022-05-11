package controllers

import (
	"fmt"
	"kuber/repo"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, repo.GetData())
}
