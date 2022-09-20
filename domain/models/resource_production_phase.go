package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resourceProductionPhase struct {
	game *Game
}

func (r resourceProductionPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) rollDices(userID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !player.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	r.game.phase = ResourceConsumption

	total := 0

	for _, dice := range r.game.dices {
		dice.number = rand.Intn(6) + 1
		total += dice.number
	}

	if total == 7 {
		r.game.phase = Robbing

		for _, player := range r.game.players {
			if len(player.resourceCards) > 7 {
				for i := 1; i <= len(player.resourceCards)/2; i++ {
					resourceCardIdx := rand.Intn(len(player.resourceCards))
					resourceCard := player.resourceCards[resourceCardIdx]
					player.resourceCards = slices.Remove(player.resourceCards, resourceCard)
					r.game.resourceCards = append(r.game.resourceCards, resourceCard)
				}
			}
		}

		return nil
	}

	dispatchedResourceCards := make([]*ResourceCard, 0)

	for _, terrain := range r.game.terrains {
		if terrain.number != total || terrain.robber != nil {
			continue
		}

		for _, player := range r.game.players {
			constructions := slices.Filter(func(construction *Construction) bool {
				return construction.land != nil && construction.land.hexCorner.isAdjacentWithHex(terrain.hex)
			}, player.constructions)

			for _, construction := range constructions {
				resourceDemand := 1
				if construction.constructionType == City {
					resourceDemand = 2
				}

				for i := 1; i <= resourceDemand; i++ {
					resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
						switch terrain.terrainType {
						case Forest:
							return resourceCard.resourceCardType == Lumber
						case Hill:
							return resourceCard.resourceCardType == Brick
						case Pasture:
							return resourceCard.resourceCardType == Wool
						case Field:
							return resourceCard.resourceCardType == Grain
						case Mountain:
							return resourceCard.resourceCardType == Ore
						}
						return false
					}, r.game.resourceCards)
					if !isExists {
						goto Rollback
					}

					r.game.resourceCards = slices.Remove(r.game.resourceCards, resourceCard)
					player.resourceCards = append(player.resourceCards, resourceCard)

					dispatchedResourceCards = append(dispatchedResourceCards, resourceCard)
				}
			}
		}
	}

	return nil

Rollback:
	for _, player := range r.game.players {
		player.resourceCards = slices.Remove(player.resourceCards, dispatchedResourceCards...)
	}
	r.game.resourceCards = append(r.game.resourceCards, dispatchedResourceCards...)

	return nil
}

func (r resourceProductionPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardID []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) maritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) offerTrading(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) confirmTrading(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceProductionPhase) cancelTrading(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}
