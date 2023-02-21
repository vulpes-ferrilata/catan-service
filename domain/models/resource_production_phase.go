package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resourceProductionPhase struct {
	game *Game
}

func (r resourceProductionPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) rollDices(userID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	r.game.phase = ResourceConsumption

	total := 0

	for _, dice := range r.game.dices {
		dice.number = rand.Intn(6) + 1
		total += dice.number
	}

	if total == 7 {
		isAnyPlayerHasResourcesExceedLimit, err := slices.Any(func(player *Player) (bool, error) {
			return len(player.resourceCards) >= 8, nil
		}, r.game.getAllPlayers()...)
		if err != nil {
			return errors.WithStack(err)
		}
		if isAnyPlayerHasResourcesExceedLimit {
			r.game.phase = ResourceDiscard
		} else {
			r.game.phase = Robbing
		}

		return nil
	}

	dispatchedResourceCards := make([]*ResourceCard, 0)

	for _, terrain := range r.game.terrains {
		if terrain.number != total || terrain.robber != nil {
			continue
		}

		for _, player := range r.game.getAllPlayers() {
			constructions, err := slices.Filter(func(construction *Construction) (bool, error) {
				isAdjacent, err := construction.land.hexCorner.isAdjacentWithHex(terrain.hex)
				if err != nil {
					return false, errors.WithStack(err)
				}
				return construction.land != nil && isAdjacent, nil
			}, player.constructions...)
			if err != nil {
				return errors.WithStack(err)
			}

			for _, construction := range constructions {
				resourceDemand := 1
				if construction.constructionType == City {
					resourceDemand = 2
				}

				for i := 1; i <= resourceDemand; i++ {
					resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
						switch terrain.terrainType {
						case Forest:
							return resourceCard.resourceCardType == Lumber, nil
						case Hill:
							return resourceCard.resourceCardType == Brick, nil
						case Pasture:
							return resourceCard.resourceCardType == Wool, nil
						case Field:
							return resourceCard.resourceCardType == Grain, nil
						case Mountain:
							return resourceCard.resourceCardType == Ore, nil
						default:
							return false, nil
						}
					}, r.game.resourceCards...)
					if errors.Is(err, slices.ErrNoMatchFound) {
						goto Rollback
					}
					if err != nil {
						return errors.WithStack(err)
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
	for _, player := range r.game.getAllPlayers() {
		player.resourceCards = slices.Remove(player.resourceCards, dispatchedResourceCards...)
	}
	r.game.resourceCards = append(r.game.resourceCards, dispatchedResourceCards...)

	return nil
}

func (r resourceProductionPhase) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardID []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}

func (r resourceProductionPhase) cancelTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceProductionPhase)
}
