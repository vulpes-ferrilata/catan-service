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
