package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
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

	nextPlayer, err := slices.Find(func(player *Player) (bool, error) {
		return player.turnOrder == r.game.activePlayer.turnOrder+1, nil
	}, r.game.players...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		r.game.turn++
		nextPlayer, err = slices.Find(func(player *Player) (bool, error) {
			return player.turnOrder == 1, nil
		}, r.game.players...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
		if err != nil {
			return errors.WithStack(err)
		}
	}
	if err != nil {
		return errors.WithStack(err)
	}

	*r.game.activePlayer, *nextPlayer = *nextPlayer, *r.game.activePlayer //swap pointer

	return nil
}

func (r resourceConsumptionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	land, err := slices.Find(func(land *Land) (bool, error) {
		return land.id == landID, nil
	}, r.game.lands...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrLandNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.useResourceCards(Lumber, Brick, Grain, Wool); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildSettlement(land); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r resourceConsumptionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	path, err := slices.Find(func(path *Path) (bool, error) {
		return path.id == pathID, nil
	}, r.game.paths...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPathNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.useResourceCards(Lumber, Brick); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildRoad(path); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r resourceConsumptionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := r.game.useResourceCards(Grain, Grain, Ore, Ore, Ore); err != nil {
		return errors.WithStack(err)
	}

	construction, err := slices.Find(func(construction *Construction) (bool, error) {
		return construction.id == constructionID, nil
	}, r.game.activePlayer.constructions...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrConstructionNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.upgradeConstruction(construction); err != nil {
		return errors.WithStack(err)
	}

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
	player, err := slices.Find(func(player *Player) (bool, error) {
		return player.userID == userID, nil
	}, r.game.getAllPlayers()...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
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
		resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
			return resourceCard.id == resourceCardID, nil
		}, player.resourceCards...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return errors.WithStack(app_errors.ErrResourceCardNotFound)
		}
		if err != nil {
			return errors.WithStack(err)
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

	requiringResourceCardQuantity, err := r.calculateRequiringResourceCardsForMaritimeTrade(specificHarborType)
	if err != nil {
		return errors.WithStack(err)
	}

	requiringResourceCards, err := slices.Filter(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.resourceCardType == resourceCardType, nil
	}, r.game.activePlayer.resourceCards...)
	if err != nil {
		return errors.WithStack(err)
	}

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

	demandingResourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.resourceCardType == demandingResourceCardType, nil
	}, r.game.resourceCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrGameHasInsufficientResourceCards)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	//add demanding resource cards
	r.game.resourceCards = slices.Remove(r.game.resourceCards, demandingResourceCard)
	r.game.activePlayer.resourceCards = append(r.game.activePlayer.resourceCards, demandingResourceCard)

	return nil
}

func (r resourceConsumptionPhase) calculateRequiringResourceCardsForMaritimeTrade(specificHarborType HarborType) (int, error) {
	isActivePlayerHasConstructionAdjacentToSpecificHarborType, err := r.isActivePlayerHasConstructionAdjacentToSpecificHarborType(specificHarborType)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	if isActivePlayerHasConstructionAdjacentToSpecificHarborType {
		return 2, nil
	}

	isActivePlayerHasConstructionAdjacentToGeneralHarborType, err := r.isActivePlayerHasConstructionAdjacentToSpecificHarborType(GeneralHarbor)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	if isActivePlayerHasConstructionAdjacentToGeneralHarborType {
		return 3, errors.WithStack(err)
	}

	return 4, nil
}

func (r resourceConsumptionPhase) isActivePlayerHasConstructionAdjacentToSpecificHarborType(harborType HarborType) (bool, error) {
	return slices.Any(func(terrain *Terrain) (bool, error) {
		if terrain.harbor == nil || terrain.harbor.harborType != harborType {
			return false, nil
		}

		intersectionHexCorners, err := findIntersectionHexCornersBetweenTwoHexes(terrain.hex, terrain.harbor.hex)
		if err != nil {
			return false, errors.WithStack(err)
		}

		return slices.Any(func(construction *Construction) (bool, error) {
			if construction.land == nil {
				return false, nil
			}

			return slices.Any(func(intersectionHexCorner HexCorner) (bool, error) {
				return construction.land.hexCorner == intersectionHexCorner, nil
			}, intersectionHexCorners...)
		}, r.game.activePlayer.constructions...)
	}, r.game.terrains...)
}

func (r resourceConsumptionPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	offeringPlayer, err := slices.Find(func(player *Player) (bool, error) {
		return player.id == playerID, nil
	}, r.game.players...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if offeringPlayer.receivedOffer {
		return errors.WithStack(app_errors.ErrYouAlreadyOfferedThisPlayer)
	}

	isActivePlayerOfferingAnyResourceCard, err := slices.Any(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.offering, nil
	}, r.game.activePlayer.resourceCards...)
	if err != nil {
		return errors.WithStack(err)
	}
	if !isActivePlayerOfferingAnyResourceCard {
		return errors.WithStack(app_errors.ErrYouMustOfferAtLeastOneResourceCard)
	}

	isOfferingPlayerSelectedAnyResourceCard, err := slices.Any(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.offering, nil
	}, offeringPlayer.resourceCards...)
	if err != nil {
		return errors.WithStack(err)
	}
	if !isOfferingPlayerSelectedAnyResourceCard {
		return errors.WithStack(app_errors.ErrSelectedPlayerMustOfferAtLeastOneResourceCard)
	}

	offeringPlayer.receivedOffer = true

	return nil
}

func (r resourceConsumptionPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	offeringPlayer, err := slices.Find(func(player *Player) (bool, error) {
		return player.userID == userID, nil
	}, r.game.players...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if !offeringPlayer.receivedOffer {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedAnyOffer)
	}

	r.game.activePlayer.receivedOffer = false
	offeringPlayer.receivedOffer = false

	selectedResourceCardsOfActivePlayer, err := slices.Filter(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.offering, nil
	}, r.game.activePlayer.resourceCards...)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(selectedResourceCardsOfActivePlayer) == 0 {
		return errors.WithStack(app_errors.ErrActivePlayerMustOfferAtLeastOneResourceCard)
	}

	for _, selectedResourceCardOfActivePlayer := range selectedResourceCardsOfActivePlayer {
		selectedResourceCardOfActivePlayer.offering = false
	}

	selectedResourceCardsOfOferringPlayer, err := slices.Filter(func(resourceCard *ResourceCard) (bool, error) {
		return resourceCard.offering, nil
	}, offeringPlayer.resourceCards...)
	if err != nil {
		return errors.WithStack(err)
	}
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
	player, err := slices.Find(func(player *Player) (bool, error) {
		return player.userID == userID, nil
	}, r.game.players...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if !player.receivedOffer {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedAnyOffer)
	}

	player.receivedOffer = false

	return nil
}
