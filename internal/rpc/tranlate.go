package rpc

import (
	"context"
	"fmt"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/server"
	"github.com/bobgo0912/bob-translate/internal/proto/translate"
	"github.com/go-redis/redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

const RedisKey = "t:%d:%s"

type TranslateRpcServer struct {
	translate.UnimplementedTranslateServer
	RedisClient *redis.Client
}

func RegService(s *server.GrpcServer, client *redis.Client) {
	s.RegService(&translate.Translate_ServiceDesc, &TranslateRpcServer{RedisClient: client})
}

func (s *TranslateRpcServer) Translate(ctx context.Context, req *translate.TranslateRequest) (*translate.TranslateReply, error) {
	if req.Lang == translate.Lang_unknown {
		return nil, status.Errorf(codes.InvalidArgument, "bad lang")
	}
	if len(req.Ids) < 1 {
		return &translate.TranslateReply{}, nil
	}
	pipeline := s.RedisClient.Pipeline()
	for u, array := range req.Ids {
		key := fmt.Sprintf(RedisKey, u, req.Lang.String())
		for _, id := range array.Ids {
			pipeline.HGet(ctx, key, fmt.Sprint(id))
		}
	}
	cmders, err := pipeline.Exec(ctx)
	if len(cmders) < 1 {
		log.Error("pipeline.Exec err=", err.Error())
		return nil, status.Errorf(codes.Internal, "redis query fail")
	}
	r := &translate.TranslateReply{
		Datas: make(map[uint64]*translate.Translated, 0),
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringCmd)
		result := cmd.Val()
		args := cmd.Args()
		arg := args[1].(string)
		id := args[2].(string)
		uid, _ := strconv.ParseUint(id, 10, 64)
		translateType := strings.Split(arg, ":")[1]
		parseUint, _ := strconv.ParseUint(translateType, 10, 64)

		if _, ok := r.Datas[parseUint]; !ok {
			r.Datas[parseUint] = &translate.Translated{Datas: make([]*translate.Id2String, 0)}
		}
		r.Datas[parseUint].Datas = append(r.Datas[parseUint].Datas, &translate.Id2String{
			Id:   uid,
			Name: result,
		})
	}
	return r, nil
}
