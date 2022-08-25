package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlayerBuilder interface {
	SetID(id primitive.ObjectID) PlayerBuilder
	SetUserID(userID primitive.ObjectID) PlayerBuilder
	SetColor(color playerColor) PlayerBuilder
	SetTurnOrder(turnOrder int) PlayerBuilder
	SetIsActive(isActive bool) PlayerBuilder
	SetIsOffered(isOffered bool) PlayerBuilder
	SetAchievements(achievements ...*Achievement) PlayerBuilder
	SetResourceCards(resourceCards ...*ResourceCard) PlayerBuilder
	SetDevelopmentCards(developmentCards ...*DevelopmentCard) PlayerBuilder
	SetConstructions(constructions ...*Construction) PlayerBuilder
	SetRoads(roads ...*Road) PlayerBuilder
	Create() *Player
}

func NewPlayerBuilder() PlayerBuilder {
	return &playerBuilder{}
}

type playerBuilder struct {
	id               primitive.ObjectID
	userID           primitive.ObjectID
	color            playerColor
	turnOrder        int
	isActive         bool
	isOffered        bool
	achievements     []*Achievement
	resourceCards    []*ResourceCard
	developmentCards []*DevelopmentCard
	constructions    []*Construction
	roads            []*Road
}

func (p *playerBuilder) SetID(id primitive.ObjectID) PlayerBuilder {
	p.id = id

	return p
}

func (p *playerBuilder) SetUserID(userID primitive.ObjectID) PlayerBuilder {
	p.userID = userID

	return p
}

func (p *playerBuilder) SetColor(color playerColor) PlayerBuilder {
	p.color = color

	return p
}

func (p *playerBuilder) SetTurnOrder(turnOrder int) PlayerBuilder {
	p.turnOrder = turnOrder

	return p
}

func (p *playerBuilder) SetIsActive(isActive bool) PlayerBuilder {
	p.isActive = isActive

	return p
}

func (p *playerBuilder) SetIsOffered(isOffered bool) PlayerBuilder {
	p.isOffered = isOffered

	return p
}

func (p *playerBuilder) SetAchievements(achievements ...*Achievement) PlayerBuilder {
	p.achievements = achievements

	return p
}

func (p *playerBuilder) SetResourceCards(resourceCards ...*ResourceCard) PlayerBuilder {
	p.resourceCards = resourceCards

	return p
}

func (p *playerBuilder) SetDevelopmentCards(developmentCards ...*DevelopmentCard) PlayerBuilder {
	p.developmentCards = developmentCards

	return p
}

func (p *playerBuilder) SetConstructions(constructions ...*Construction) PlayerBuilder {
	p.constructions = constructions

	return p
}

func (p *playerBuilder) SetRoads(roads ...*Road) PlayerBuilder {
	p.roads = roads

	return p
}

func (p playerBuilder) Create() *Player {
	return &Player{
		aggregate: aggregate{
			id: p.id,
		},
		userID:           p.userID,
		color:            p.color,
		turnOrder:        p.turnOrder,
		isOffered:        p.isOffered,
		isActive:         p.isActive,
		achievements:     p.achievements,
		resourceCards:    p.resourceCards,
		developmentCards: p.developmentCards,
		constructions:    p.constructions,
		roads:            p.roads,
	}
}
