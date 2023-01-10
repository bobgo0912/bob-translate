package translate

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestR(t *testing.T) {
	//stream
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := NewTranslateClient(conn)
	r, err := c.Translate(context.Background(), &TranslateRequest{
		Lang: Lang_zh,
		Ids: map[uint64]*Uint64Array{1: &Uint64Array{Ids: []uint64{1, 2, 3}},
			2: &Uint64Array{Ids: []uint64{1, 2}},
			3: &Uint64Array{Ids: []uint64{1, 2, 3, 4}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r.Message)
}
