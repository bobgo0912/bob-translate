package util

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/bob-translate/internal/proto/translate"
	"github.com/pkg/errors"
)

type Market uint64

const (
	NameMarket   Market = 1
	MarketMarket Market = 2
	TeamMarket   Market = 3
)

type Id2Str struct {
	Id  uint64
	Str *string
}

type TranslateTool struct {
	Data map[uint64][]*Id2Str
	req  map[uint64]*translate.Uint64Array
}

func NewTranslate() *TranslateTool {
	return &TranslateTool{Data: map[uint64][]*Id2Str{}, req: map[uint64]*translate.Uint64Array{}}
}
func (t *TranslateTool) AddName(id uint64, str *string) {
	_, ok := t.Data[uint64(NameMarket)]
	if !ok {
		t.Data[uint64(NameMarket)] = make([]*Id2Str, 0)
	}
	t.Data[uint64(NameMarket)] = append(t.Data[uint64(NameMarket)], &Id2Str{
		Id:  id,
		Str: str,
	})
	_, ok = t.req[uint64(NameMarket)]
	if !ok {
		t.req[uint64(NameMarket)] = &translate.Uint64Array{Ids: make([]uint64, 0)}
	}
	t.req[uint64(NameMarket)].Ids = append(t.req[uint64(NameMarket)].Ids, id)
}
func (t *TranslateTool) AddMarket(id uint64, str *string) {
	_, ok := t.Data[uint64(MarketMarket)]
	if !ok {
		t.Data[uint64(MarketMarket)] = make([]*Id2Str, 0)
	}
	t.Data[uint64(MarketMarket)] = append(t.Data[uint64(MarketMarket)], &Id2Str{
		Id:  id,
		Str: str,
	})
	_, ok = t.req[uint64(MarketMarket)]
	if !ok {
		t.req[uint64(MarketMarket)] = &translate.Uint64Array{Ids: make([]uint64, 0)}
	}
	t.req[uint64(MarketMarket)].Ids = append(t.req[uint64(MarketMarket)].Ids, id)
}
func (t *TranslateTool) AddTeam(id uint64, str *string) {
	_, ok := t.Data[uint64(TeamMarket)]
	if !ok {
		t.Data[uint64(TeamMarket)] = make([]*Id2Str, 0)
	}
	t.Data[uint64(TeamMarket)] = append(t.Data[uint64(TeamMarket)], &Id2Str{
		Id:  id,
		Str: str,
	})
	_, ok = t.req[uint64(TeamMarket)]
	if !ok {
		t.req[uint64(TeamMarket)] = &translate.Uint64Array{Ids: make([]uint64, 0)}
	}
	t.req[uint64(TeamMarket)].Ids = append(t.req[uint64(TeamMarket)].Ids, id)
}
func (t *TranslateTool) Translate(ctx context.Context, client translate.TranslateClient, lan translate.Lang) error {
	reply, err := client.Translate(ctx, &translate.TranslateRequest{Lang: lan, Ids: t.req})
	if err != nil {
		log.Error("Translate fail err=", err.Error())
		return errors.Wrap(err, "Translate fail")
	}
	for u, translated := range reply.Datas {
		for _, data := range translated.Datas {
			for _, str := range t.Data[u] {
				if data.Id == str.Id {
					*str.Str = data.Name
					break
				}
			}
		}
	}
	return nil
}
