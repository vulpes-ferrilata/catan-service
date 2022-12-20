package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlayerBuilder struct {
	id                 primitive.ObjectID
	userID             primitive.ObjectID
	color              PlayerColor
	turnOrder          int
	receivedOffer      bool
	discardedResources bool
	score              int
	achievements       []*Achievement
	resourceCards      []*ResourceCard
	developmentCards   []*DevelopmentCard
	constructions      []*Construction
	roads              []*Road
}

func (p PlayerBuilder) SetID(id primitive.ObjectID) PlayerBuilder {
	p.id = id

	return p
}

func (p PlayerBuilder) SetUserID(userID primitive.ObjectID) PlayerBuilder {
	p.userID = userID

	return p
}

func (p PlayerBuilder) SetColor(color PlayerColor) PlayerBuilder {
	p.color = color

	return p
}

func (p PlayerBuilder) SetTurnOrder(turnOrder int) PlayerBuilder {
	p.turnOrder = turnOrder

	return p
}

func (p PlayerBuilder) SetReceivedOffer(receivedOffer bool) PlayerBuilder {
	p.receivedOffer = receivedOffer

	return p
}

func (p PlayerBuilder) SetDiscardedResources(discardedResources bool) PlayerBuilder {
	p.discardedResources = discardedResources

	return p
}

func (p PlayerBuilder) SetScore(score int) PlayerBuilder {
	p.score = score

	return p
}

func (p PlayerBuilder) SetAchievements(achievements ...*Achievement) PlayerBuilder {
	p.achievements = achievements

	return p
}

func (p PlayerBuilder) SetResourceCards(resourceCards ...*ResourceCard) PlayerBuilder {
	p.resourceCards = resourceCards

	return p
}

func (p PlayerBuilder) SetDevelopmentCards(developmentCards ...*DevelopmentCard) PlayerBuilder {
	p.developmentCards = developmentCards

	return p
}

func (p PlayerBuilder) SetConstructions(constructions ...*Construction) PlayerBuilder {
	p.constructions = constructions

	return p
}

func (p PlayerBuilder) SetRoads(roads ...*Road) PlayerBuilder {
	p.roads = roads

	return p
}

func (p PlayerBuilder) Create() *Player {
	return &Player{
		aggregate: aggregate{
			id: p.id,
		},
		userID:             p.userID,
		color:              p.color,
		turnOrder:          p.turnOrder,
		receivedOffer:      p.receivedOffer,
		discardedResources: p.discardedResources,
		score:              p.score,
		achievements:       p.achievements,
		resourceCards:      p.resourceCards,
		developmentCards:   p.developmentCards,
		constructions:      p.constructions,
		roads:              p.roads,
	}
}
