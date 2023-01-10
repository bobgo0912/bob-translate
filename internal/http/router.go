package http

import (
	"fmt"
	"github.com/bobgo0912/b0b-common/pkg/log"
	hello "github.com/bobgo0912/b0b-common/pkg/server/proto"
	"github.com/bobgo0912/bob-translate/internal/service"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type TranslateRouter struct {
	Service *service.TranslateService
}

func (r *TranslateRouter) Translate(req proto.Message, w http.ResponseWriter) {
	log.Info(req)
	request := req.(*hello.HelloRequest)
	fmt.Println(request)
	w.WriteHeader(http.StatusOK)
	reply := hello.HelloReply{Message: "drrr"}
	marshal, _ := proto.Marshal(&reply)
	w.Write(marshal)
}
