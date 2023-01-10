package util

import (
	"context"
	"fmt"
	"github.com/bobgo0912/bob-translate/internal/proto/translate"
	"google.golang.org/grpc"
	"testing"
)

type Test struct {
	Id   uint64
	Name string
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
	tests := make([]*Test, 0)
	for i := 0; i < 4; i++ {
		var sd Test
		sd.Id = uint64(i + 1)
		tests = append(tests, &sd)
	}
	for _, test := range tests {
		fmt.Printf("%x\n", &test.Name)
		translateTool.AddName(test.Id, &test.Name)
	}

	idd := make([]*Test, 0)
	for i := 0; i < 3; i++ {
		var sd Test
		sd.Id = uint64(i + 1)
		idd = append(idd, &sd)
	}

	for _, test := range idd {
		fmt.Printf("%x\n", &test.Name)
		translateTool.AddTeam(test.Id, &test.Name)
	}

	err = translateTool.Translate(context.Background(), c, translate.Lang_zh)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(tests)
}
