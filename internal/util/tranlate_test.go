package util

import (
	"context"
	"github.com/bobgo0912/bob-translate/internal/proto/translate"
	"google.golang.org/grpc"
	"testing"
)

type Test struct {
	Id   uint64
	Name *string
}

func TestTranslate(t *testing.T) {
	//stream
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := translate.NewTranslateClient(conn)
	translateTool := NewTranslate()
	tests := []Test{{
		Id: 1, Name: new(string),
	}, {Id: 2, Name: new(string)},
		{Id: 3, Name: new(string)}}
	for _, test := range tests {
		translateTool.AddName(test.Id, test.Name)
	}
	err = translateTool.Translate(context.Background(), c, translate.Lang_zh)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tests)
}
