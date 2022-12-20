package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resourceConsumptionPhase struct {
	game *Game
}

func (r resourceConsumptionPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceConsumptionPhase)
}

func (r resourceConsumptionPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceConsumptionPhase)
}

func (r resourceConsumptionPhase) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceConsumptionPhase)
}

func (r resourceConsumptionPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceConsumptionPhase)
}

func (r resourceConsumptionPhase) endTurn(userID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	r.game.phase = ResourceProduction

	for _, player := range r.game.getAllPlayers() {
		player.discardedResources = false

		for _, resourceCard := range player.resourceCards {
			resourceCard.offering = false
		}

		for _, developmentCard := range player.developmentCards {
			if developmentCard.status == Disable {
				developmentCard.status = Enable
			}
		}
	}

	nextPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.turnOrder == r.game.activePlayer.turnOrder+1
	}, r.game.players)
	if !isExists {
		r.game.turn++
		nextPlayer, isExists = slices.Find(func(player *Player) bool {
			return player.turnOrder == 1
		}, r.game.players)
		if !isExists {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
	}

	*r.game.activePlayer, *nextPlayer = *nextPlayer, *r.game.activePlayer //swap pointer

	return nil
}

func (r resourceConsumptionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	land, isExists := slices.Find(func(land *Land) bool {
		return land.id == landID
	}, r.game.lands)
	if !isExists {
		return errors.WithStack(app_errors.ErrLandNotFound)
	}

	if err := r.game.useResourceCards(Lumber, Brick, Grain, Wool); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildSettlement(land); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	path, isExists := slices.Find(func(path *Path) bool {
		return path.id == pathID
	}, r.game.paths)
	if !isExists {
		return errors.WithStack(app_errors.ErrPathNotFound)
	}

	if err := r.game.useResourceCards(Lumber, Brick); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildRoad(path); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := r.game.useResourceCards(Grain, Grain, Ore, Ore, Ore); err != nil {
		return errors.WithStack(err)
	}

	construction, isExists := slices.Find(func(construction *Construction) bool {
		return construction.id == constructionID
	}, r.game.activePlayer.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrConstructionNotFound)
	}

	if err := r.game.upgradeConstruction(construction); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := r.game.useResourceCards(Wool, Grain, Ore); err != nil {
		return errors.WithStack(err)
	}

	if len(r.game.developmentCards) == 0 {
		return errors.WithStack(app_errors.ErrGameRunOutOfDevelopmentCards)
	}

	developmentCardIdx := rand.Intn(len(r.game.developmentCards))
	developmentCard := r.game.developmentCards[developmentCardIdx]

	if developmentCard.isVictoryPointCard() {
		developmentCard.status = Enable
	}

	r.game.developmentCards = slices.Remove(r.game.developmentCards, developmentCard)
	r.game.activePlayer.developmentCards = append(r.game.activePlayer.developmentCards, developmentCard)

	return nil
}

func (r resourceConsumptionPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.getAllPlayers())
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	//cancel offer of all players
	if player == r.game.activePlayer {
		for _, player := range r.game.players {
			player.receivedOffer = false
		}
	} else {
		player.receivedOffer = false
	}

	for _, resourceCardID := range resourceCardIDs {
		resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
			return resourceCard.id == resourceCardID
		}, player.resourceCards)
		if !isExists {
			return errors.WithStack(app_errors.ErrResourceCardNotFound)
		}

		resourceCard.offering = !resourceCard.offering
	}

	return nil
}

func (r resourceConsumptionPhase) maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	var requiringResourceCardQuantity int

	var specificHarborType HarborType

	switch resourceCardType {
	case Lumber:
		specificHarborType = LumberHarbor
	case Brick:
		specificHarborType = BrickHarbor
	case Wool:
		specificHarborType = WoolHarbor
	case Grain:
		specificHarborType = GrainHarbor
	case Ore:
		specificHarborType = OreHarbor
	}

	if r.isActivePlayerHasConstructionAdjacentToSpecificHarborType(specificHarborType) {
		requiringResourceCardQuantity = 2
	} else if r.isActivePlayerHasConstructionAdjacentToSpecificHarborType(GeneralHarbor) {
		requiringResourceCardQuantity = 3
	} else {
		requiringResourceCardQuantity = 4
	}

	requiringResourceCards := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.resourceCardType == resourceCardType
	}, r.game.activePlayer.resourceCards)

	if len(requiringResourceCards) < requiringResourceCardQuantity {
		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
	}

	requiringResourceCards = requiringResourceCards[:requiringResourceCardQuantity]

	for _, requiringResourceCard := range requiringResourceCards {
		requiringResourceCard.offering = false
	}

	//substract requiring resource cards
	r.game.activePlayer.resourceCards = slices.Remove(r.game.activePlayer.resourceCards, requiringResourceCards...)
	r.game.resourceCards = append(r.game.resourceCards, requiringResourceCards...)

	demandingResourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
		return resourceCard.resourceCardType == demandingResourceCardType
	}, r.game.resourceCards)
	if !isExists {
		return errors.WithStack(app_errors.ErrGameHasInsufficientResourceCards)
	}

	//add demanding resource cards
	r.game.resourceCards = slices.Remove(r.game.resourceCards, demandingResourceCard)
	r.game.activePlayer.resourceCards = append(r.game.activePlayer.resourceCards, demandingResourceCard)

	return nil
}

func (r resourceConsumptionPhase) isActivePlayerHasConstructionAdjacentToSpecificHarborType(harborType HarborType) bool {
	return slices.Any(func(terrain *Terrain) bool {
		if terrain.harbor == nil || terrain.harbor.harborType != harborType {
			return false
		}

		intersectionHexCorners := findIntersectionHexCornersBetweenTwoHexes(terrain.hex, terrain.harbor.hex)

		return slices.Any(func(construction *Construction) bool {
			if construction.land == nil {
				return false
			}

			return slices.Any(func(intersectionHexCorner HexCorner) bool {
				return construction.land.hexCorner == intersectionHexCorner
			}, intersectionHexCorners)
		}, r.game.activePlayer.constructions)
	}, r.game.terrains)
}

func (r resourceConsumptionPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	offeringPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.id == playerID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if offeringPlayer.receivedOffer {
		return errors.WithStack(app_errors.ErrYouAlreadyOfferedThisPlayer)
	}

	isActivePlayerSelectedAnyResourceCard := slices.Any(func(resourceCard *ResourceCard) bool {
		return resourceCard.offering
	}, r.game.activePlayer.resourceCards)
	if !isActivePlayerSelectedAnyResourceCard {
		return errors.WithStack(app_errors.ErrYouMustOfferAtLeastOneResourceCard)
	}

	isOfferingPlayerSelectedAnyResourceCard := slices.Any(func(resourceCard *ResourceCard) bool {
		return resourceCard.offering
	}, offeringPlayer.resourceCards)
	if !isOfferingPlayerSelectedAnyResourceCard {
		return errors.WithStack(app_errors.ErrSelectedPlayerMustOfferAtLeastOneResourceCard)
	}

	offeringPlayer.receivedOffer = true

	return nil
}

func (r resourceConsumptionPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	offeringPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if !offeringPlayer.receivedOffer {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedAnyOffer)
	}

	r.game.activePlayer.receivedOffer = false
	offeringPlayer.receivedOffer = false

	selectedResourceCardsOfActivePlayer := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.offering
	}, r.game.activePlayer.resourceCards)
	if len(selectedResourceCardsOfActivePlayer) == 0 {
		return errors.WithStack(app_errors.ErrActivePlayerMustOfferAtLeastOneResourceCard)
	}

	for _, selectedResourceCardOfActivePlayer := range selectedResourceCardsOfActivePlayer {
		selectedResourceCardOfActivePlayer.offering = false
	}

	selectedResourceCardsOfOferringPlayer := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.offering
	}, offeringPlayer.resourceCards)
	if len(selectedResourceCardsOfOferringPlayer) == 0 {
		return errors.WithStack(app_errors.ErrYouMustOfferAtLeastOneResourceCard)
	}

	for _, selectedResourceCardOfOferringPlayer := range selectedResourceCardsOfOferringPlayer {
		selectedResourceCardOfOferringPlayer.offering = false
	}
	//swap resource cards
	r.game.activePlayer.resourceCards = slices.Remove(r.game.activePlayer.resourceCards, selectedResourceCardsOfActivePlayer...)
	offeringPlayer.resourceCards = append(offeringPlayer.resourceCards, selectedResourceCardsOfActivePlayer...)

	offeringPlayer.resourceCards = slices.Remove(offeringPlayer.resourceCards, selectedResourceCardsOfOferringPlayer...)
	r.game.activePlayer.resourceCards = append(r.game.activePlayer.resourceCards, selectedResourceCardsOfOferringPlayer...)

	return nil
}

func (r resourceConsumptionPhase) cancelTradeOffer(userID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if !player.receivedOffer {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedAnyOffer)
	}

	player.receivedOffer = false

	return nil
}
