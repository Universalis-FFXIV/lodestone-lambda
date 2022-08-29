package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

type LodestoneCharacterEvent struct {
	ID uint32 `json:"id"`
}

type LodestoneCharacterResult struct {
	Bio    string `json:"bio"`
	Name   string `json:"name"`
	World  string `json:"world"`
	Avatar string `json:"avatar"`
}

func HandleRequest(ctx context.Context, e LodestoneCharacterEvent) (*LodestoneCharacterResult, error) {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	character, err := s.FetchCharacter(e.ID)
	if err != nil {
		return nil, err
	}

	res := LodestoneCharacterResult{
		Bio:    character.Bio,
		Name:   character.Name,
		World:  character.World,
		Avatar: character.Avatar,
	}

	return &res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
