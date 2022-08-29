package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

type LodestoneSearchEvent struct {
	World     string `json:"world"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LodestoneSearchResult struct {
	ID uint32 `json:"id"`
}

func HandleRequest(ctx context.Context, e LodestoneSearchEvent) (*LodestoneSearchResult, error) {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	worldName := strings.ToLower(e.World)
	if worldName == "" {
		return nil, errors.New("world name not provided")
	}

	characterName := strings.ToLower(fmt.Sprintf("%s %s", e.FirstName, e.LastName))
	if characterName == "" {
		return nil, errors.New("character name not provided")
	}

	for res := range s.SearchCharacters(godestone.CharacterOptions{
		Name:  characterName,
		World: strings.ToUpper(string(worldName[0])) + worldName[1:], // World name must be captialized
	}) {
		if res.Error != nil {
			return nil, res.Error
		}

		if strings.ToLower(res.Name) == characterName && strings.ToLower(res.World) == worldName {
			r := LodestoneSearchResult{
				ID: res.ID,
			}

			return &r, nil
		}
	}

	return nil, errors.New("no character matching those parameters was found")
}

func main() {
	lambda.Start(HandleRequest)
}
