package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

type LodestoneAchievementsEvent struct {
	ID string `json:"id"`
}

type LodestoneAchievementsResult struct {
	Achievements []uint32 `json:"achievements"`
	Private      bool     `json:"private"`
}

func HandleRequest(ctx context.Context, e LodestoneAchievementsEvent) (*LodestoneAchievementsResult, error) {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	characterId, err := strconv.ParseUint(e.ID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("[BadRequest] %s", err.Error())
	}

	achievements, allInfo, err := s.FetchCharacterAchievements(uint32(characterId))
	if err != nil {
		return nil, fmt.Errorf("[InternalServerError] %s", err.Error())
	}

	if allInfo.Private {
		return &LodestoneAchievementsResult{
			Achievements: []uint32{},
			Private:      true,
		}, nil
	}

	ids := make([]uint32, 0, len(achievements))
	for _, a := range achievements {
		ids = append(ids, a.ID)
	}

	return &LodestoneAchievementsResult{
		Achievements: ids,
		Private:      false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
