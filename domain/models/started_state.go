package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type startedState struct {
	game *Game
}

func (s startedState) getPhase() phase {
	switch s.game.phase {
	case Setup:
		return &setupPhase{s.game}
	case ResourceProduction:
		return &resourceProductionPhase{s.game}
	case ResourceDiscard:
		return &resourceDiscardPhase{s.game}
	case Robbing:
		return &robbingPhase{s.game}
	case ResourceConsumption:
		return &resourceConsumptionPhase{s.game}
	}

	return nil
}

func (s startedState) newPlayer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyStarted)
}

func (s startedState) startGame(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyStarted)
}

func (s startedState) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := s.getPhase().buildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) rollDices(userID primitive.ObjectID) error {
	if err := s.getPhase().rollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := s.getPhase().discardResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := s.getPhase().moveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) endTurn(userID primitive.ObjectID) error {
	if err := s.getPhase().endTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	if err := s.getPhase().buildSettlement(userID, landID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := s.getPhase().buildRoad(userID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if err := s.getPhase().upgradeCity(userID, constructionID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buyDevelopmentCard(userID primitive.ObjectID) error {
	if err := s.getPhase().buyDevelopmentCard(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := s.getPhase().toggleResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	if err := s.getPhase().maritimeTrade(userID, resourceCardType, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := s.getPhase().sendTradeOffer(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) confirmTradeOffer(userID primitive.ObjectID) error {
	if err := s.getPhase().confirmTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) cancelTradeOffer(userID primitive.ObjectID) error {
	if err := s.getPhase().cancelTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) playKnightCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	developmentCard, err := slices.Find(func(developmentCard *DevelopmentCard) (bool, error) {
		return developmentCard.id == developmentCardID, nil
	}, s.game.activePlayer.developmentCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if developmentCard.developmentCardType != Knight {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardMustBeKnightCard)
	}

	if developmentCard.status != Enable {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardIsUnavailableToUse)
	}

	developmentCard.status = Used

	for _, developmentCard := range s.game.activePlayer.developmentCards {
		if developmentCard.status == Enable && !developmentCard.isVictoryPointCard() {
			developmentCard.status = Disable
		}
	}

	terrain, err := slices.Find(func(terrain *Terrain) (bool, error) {
		return terrain.id == terrainID, nil
	}, s.game.terrains...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	var player *Player
	if playerID != primitive.NilObjectID {
		player, err = slices.Find(func(player *Player) (bool, error) {
			return player.id == playerID, nil
		}, s.game.players...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if err := s.game.moveRobber(terrain); err != nil {
		return errors.WithStack(err)
	}

	if err := s.game.robPlayer(player); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) playRoadBuildingCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, pathIDs []primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	developmentCard, err := slices.Find(func(developmentCard *DevelopmentCard) (bool, error) {
		return developmentCard.id == developmentCardID, nil
	}, s.game.activePlayer.developmentCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if developmentCard.developmentCardType != RoadBuilding {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardMustBeRoadBuildingCard)
	}

	if developmentCard.status != Enable {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardIsUnavailableToUse)
	}

	developmentCard.status = Used

	for _, developmentCard := range s.game.activePlayer.developmentCards {
		if developmentCard.status == Enable && !developmentCard.isVictoryPointCard() {
			developmentCard.status = Disable
		}
	}

	paths, err := slices.Map(func(pathID primitive.ObjectID) (*Path, error) {
		path, err := slices.Find(func(path *Path) (bool, error) {
			return path.id == pathID, nil
		}, s.game.paths...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return nil, errors.WithStack(app_errors.ErrPathNotFound)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return path, nil
	}, pathIDs...)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, path := range paths {
		if err := s.game.buildRoad(path); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (s startedState) playYearOfPlentyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, demandingResourceCardTypes []ResourceCardType) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	developmentCard, err := slices.Find(func(developmentCard *DevelopmentCard) (bool, error) {
		return developmentCard.id == developmentCardID, nil
	}, s.game.activePlayer.developmentCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if developmentCard.developmentCardType != YearOfPlenty {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardMustBeYearOfPlentyCard)
	}

	if developmentCard.status != Enable {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardIsUnavailableToUse)
	}

	developmentCard.status = Used

	for _, developmentCard := range s.game.activePlayer.developmentCards {
		if developmentCard.status == Enable && !developmentCard.isVictoryPointCard() {
			developmentCard.status = Disable
		}
	}

	if len(demandingResourceCardTypes) == 1 {
		demandingResourceCardTypes = append(demandingResourceCardTypes, demandingResourceCardTypes...)
	}

	for _, resourceCardType := range demandingResourceCardTypes {
		resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
			return resourceCard.resourceCardType == resourceCardType, nil
		}, s.game.resourceCards...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			continue
		}
		if err != nil {
			return errors.WithStack(err)
		}

		s.game.resourceCards = slices.Remove(s.game.resourceCards, resourceCard)
		s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, resourceCard)
	}

	return nil
}

func (s startedState) playMonopolyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	developmentCard, err := slices.Find(func(developmentCard *DevelopmentCard) (bool, error) {
		return developmentCard.id == developmentCardID, nil
	}, s.game.activePlayer.developmentCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if developmentCard.developmentCardType != Monopoly {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardMustBeMonopolyCard)
	}

	if developmentCard.status != Enable {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardIsUnavailableToUse)
	}

	developmentCard.status = Used

	for _, developmentCard := range s.game.activePlayer.developmentCards {
		if developmentCard.status == Enable && !developmentCard.isVictoryPointCard() {
			developmentCard.status = Disable
		}
	}

	demandingResourceCards := make([]*ResourceCard, 0)

	for _, player := range s.game.players {
		for _, resourceCard := range player.resourceCards {
			if resourceCard.resourceCardType == demandingResourceCardType {
				player.resourceCards = slices.Remove(player.resourceCards, resourceCard)
				demandingResourceCards = append(demandingResourceCards, resourceCard)
			}
		}
	}

	s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, demandingResourceCards...)

	return nil
}

func (s startedState) playVictoryPointCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	developmentCard, err := slices.Find(func(developmentCard *DevelopmentCard) (bool, error) {
		return developmentCard.id == developmentCardID, nil
	}, s.game.activePlayer.developmentCards...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}

	if !developmentCard.isVictoryPointCard() {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardMustBeVictoryPointCard)
	}

	if developmentCard.status != Enable {
		return errors.WithStack(app_errors.ErrSelectedDevelopmentCardIsUnavailableToUse)
	}

	developmentCard.status = Used

	return nil
}
