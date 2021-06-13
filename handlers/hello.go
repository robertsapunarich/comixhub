package handlers

import (
	"comixhub/utils"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	utils.RenderJSON(writer, "hello!")
}
