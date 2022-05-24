package controllers

import (
	"fmt"
	"kuber/repo"
	"net/http"
)

func GetHash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, repo.HashOfFile("/home/tshorokhov@scnsoft.com/Pictures/1"))
}
