package models

type DevelopmentCard struct {
	aggregate
	developmentCardType DevelopmentCardType
	status              developmentCardStatus
}

func (r DevelopmentCard) GetType() DevelopmentCardType {
	return r.developmentCardType
}

func (r DevelopmentCard) GetStatus() developmentCardStatus {
	return r.status
}

func (r DevelopmentCard) isVictoryPointCard() bool {
	switch r.developmentCardType {
	case Chapel, GreatHall, Library, Market, University:
		return true
	default:
		return false
	}
}
