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
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceConsumptionPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceConsumptionPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceConsumptionPhase) endTurn(userID primitive.ObjectID) error {
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !activePlayer.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	r.game.phase = ResourceProduction

	for _, player := range r.game.players {
		for _, resourceCard := range player.resourceCards {
			resourceCard.isSelected = false
		}

		for _, developmentCard := range player.developmentCards {
			if developmentCard.status == Disable {
				developmentCard.status = Enable
			}
		}

		player.isActive = false
	}

	nextPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.turnOrder == activePlayer.turnOrder+1
	}, r.game.players)
	if !isExists {
		r.game.turn++
		nextPlayer, isExists := slices.Find(func(player *Player) bool {
			return player.turnOrder == 1
		}, r.game.players)
		if !isExists {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
		nextPlayer.isActive = true
		return nil
	}
	nextPlayer.isActive = true

	return nil
}

func (r resourceConsumptionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	land, isExists := slices.Find(func(land *Land) bool {
		return land.id == landID
	}, r.game.lands)
	if !isExists {
		return errors.WithStack(app_errors.ErrLandNotFound)
	}

	if err := r.game.useResourceCards(player, Lumber, Brick, Grain, Wool); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildSettlement(player, land); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	path, isExists := slices.Find(func(path *Path) bool {
		return path.id == pathID
	}, r.game.paths)
	if !isExists {
		return errors.WithStack(app_errors.ErrPathNotFound)
	}

	if err := r.game.useResourceCards(player, Lumber, Brick); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.buildRoad(player, path); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if err := r.game.useResourceCards(player, Grain, Grain, Ore, Ore, Ore); err != nil {
		return errors.WithStack(err)
	}

	construction, isExists := slices.Find(func(construction *Construction) bool {
		return construction.id == constructionID
	}, player.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrConstructionNotFound)
	}

	if err := r.game.upgradeConstruction(player, construction); err != nil {
		return errors.WithStack(err)
	}

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if err := r.game.useResourceCards(player, Wool, Grain, Ore); err != nil {
		return errors.WithStack(err)
	}

	if len(r.game.developmentCards) == 0 {
		return errors.WithStack(app_errors.ErrGameHasInsufficientDevelopmentCards)
	}

	developmentCardIdx := rand.Intn(len(r.game.developmentCards))
	developmentCard := r.game.developmentCards[developmentCardIdx]
	r.game.developmentCards = slices.Remove(r.game.developmentCards, developmentCard)
	player.developmentCards = append(player.developmentCards, developmentCard)

	r.game.calculateScore()

	return nil
}

func (r resourceConsumptionPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	player.isOffered = false

	if player.isActive {
		for _, player := range r.game.players {
			player.isOffered = false
		}
	}

	for _, resourceCardID := range resourceCardIDs {
		resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
			return resourceCard.id == resourceCardID
		}, player.resourceCards)
		if !isExists {
			return errors.WithStack(app_errors.ErrResourceCardNotFound)
		}

		resourceCard.isSelected = !resourceCard.isSelected
	}

	return nil
}

func (r resourceConsumptionPhase) maritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !player.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	totalDemandingResourceCardQuantity := 0

	requiringResourceCardTypes := []ResourceCardType{
		Lumber,
		Brick,
		Wool,
		Grain,
		Ore,
	}

	requiringResourceCardTypes = slices.Remove(requiringResourceCardTypes, demandingResourceCardType)

	for _, requiringResourceCardType := range requiringResourceCardTypes {
		requiringResourceCardQuantity := 4

		var specificHarborType HarborType

		switch requiringResourceCardType {
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

		if r.isPlayerHasConstructionAdjacentToSpecificHarborType(player, specificHarborType) {
			requiringResourceCardQuantity = 2
		} else if r.isPlayerHasConstructionAdjacentToSpecificHarborType(player, GeneralHarbor) {
			requiringResourceCardQuantity = 3
		}

		requiringResourceCards := slices.Filter(func(resourceCard *ResourceCard) bool {
			return resourceCard.resourceCardType == requiringResourceCardType && resourceCard.isSelected
		}, player.resourceCards)

		demandingResourceCardQuantity := len(requiringResourceCards) / requiringResourceCardQuantity

		requiringResourceCards = requiringResourceCards[:demandingResourceCardQuantity*requiringResourceCardQuantity]

		for _, requiringResourceCard := range requiringResourceCards {
			requiringResourceCard.isSelected = false
		}
		//substract requiring resource cards
		player.resourceCards = slices.Remove(player.resourceCards, requiringResourceCards...)
		r.game.resourceCards = append(r.game.resourceCards, requiringResourceCards...)

		totalDemandingResourceCardQuantity += demandingResourceCardQuantity
	}

	demandingResourceCards := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.resourceCardType == demandingResourceCardType
	}, r.game.resourceCards)

	if len(demandingResourceCards) < totalDemandingResourceCardQuantity {
		return errors.WithStack(app_errors.ErrGameHasInsufficientResourceCards)
	}

	demandingResourceCards = demandingResourceCards[:totalDemandingResourceCardQuantity]
	//add demanding resource cards
	r.game.resourceCards = slices.Remove(r.game.resourceCards, demandingResourceCards...)
	player.resourceCards = append(player.resourceCards, demandingResourceCards...)

	return nil
}

func (r resourceConsumptionPhase) isPlayerHasConstructionAdjacentToSpecificHarborType(player *Player, harborType HarborType) bool {
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
		}, player.constructions)
	}, r.game.terrains)
}

func (r resourceConsumptionPhase) offerTrading(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !player.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	offeringPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.id == playerID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if offeringPlayer == player {
		return errors.WithStack(app_errors.ErrYouCannotOfferYourself)
	}

	if offeringPlayer.isOffered {
		return errors.WithStack(app_errors.ErrYouAlreadyOfferedThisPlayer)
	}

	isPlayerSelectedAnyResourceCard := slices.Any(func(resourceCard *ResourceCard) bool {
		return resourceCard.isSelected
	}, player.resourceCards)

	if !isPlayerSelectedAnyResourceCard {
		return errors.WithStack(app_errors.ErrYouMustSelectAtLeastOneResourceCard)
	}

	isOfferingPlayerSelectedAnyResourceCard := slices.Any(func(resourceCard *ResourceCard) bool {
		return resourceCard.isSelected
	}, offeringPlayer.resourceCards)

	if !isOfferingPlayerSelectedAnyResourceCard {
		return errors.WithStack(app_errors.ErrSelectedPlayerMustSelectAtLeastOneResourceCard)
	}

	offeringPlayer.isOffered = true

	return nil
}

func (r resourceConsumptionPhase) confirmTrading(userID primitive.ObjectID) error {
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.isActive
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	offeringPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if offeringPlayer == activePlayer {
		return errors.WithStack(app_errors.ErrYouCannotTradeWithYourself)
	}

	if !offeringPlayer.isOffered {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedOffer)
	}

	activePlayer.isOffered = false
	offeringPlayer.isOffered = false

	selectedResourceCardsOfActivePlayer := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.isSelected
	}, activePlayer.resourceCards)

	if len(selectedResourceCardsOfActivePlayer) == 0 {
		return errors.WithStack(app_errors.ErrActivePlayerMustSelectAtLeastOneResourceCard)
	}

	for _, selectedResourceCardOfActivePlayer := range selectedResourceCardsOfActivePlayer {
		selectedResourceCardOfActivePlayer.isSelected = false
	}

	selectedResourceCardsOfOferringPlayer := slices.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.isSelected
	}, offeringPlayer.resourceCards)

	if len(selectedResourceCardsOfOferringPlayer) == 0 {
		return errors.WithStack(app_errors.ErrYouMustSelectAtLeastOneResourceCard)
	}

	for _, selectedResourceCardOfOferringPlayer := range selectedResourceCardsOfOferringPlayer {
		selectedResourceCardOfOferringPlayer.isSelected = false
	}

	activePlayer.resourceCards = slices.Remove(activePlayer.resourceCards, selectedResourceCardsOfActivePlayer...)
	offeringPlayer.resourceCards = append(offeringPlayer.resourceCards, selectedResourceCardsOfActivePlayer...)

	offeringPlayer.resourceCards = slices.Remove(offeringPlayer.resourceCards, selectedResourceCardsOfOferringPlayer...)
	activePlayer.resourceCards = append(activePlayer.resourceCards, selectedResourceCardsOfOferringPlayer...)

	return nil
}

func (r resourceConsumptionPhase) cancelTrading(userID primitive.ObjectID) error {
	player, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if !player.isOffered {
		return errors.WithStack(app_errors.ErrYouHaveNotReceivedOffer)
	}

	player.isOffered = false

	return nil
}
