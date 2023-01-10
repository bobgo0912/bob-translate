package http

import (
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/server"
	hello "github.com/bobgo0912/b0b-common/pkg/server/proto"
	"github.com/bobgo0912/bob-translate/internal/service"
	"net/http"
)

type TranslateMain struct {
	R *TranslateRouter
}

func NewTranslateMain(r *server.MuxRouter) *TranslateMain {
	router := &TranslateRouter{Service: &service.TranslateService{}}
	main := TranslateMain{R: router}
	r.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("test")
		writer.Write([]byte("ttt"))
	}).Methods("GET")
	r.HandleProtoFunc("/proto", router.Translate, &hello.HelloRequest{}).Methods("POST", "OPTIONS")
	return &main
}
