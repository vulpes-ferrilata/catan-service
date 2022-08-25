package models

type DevelopmentCard struct {
	aggregate
	developmentCardType developmentCardType
	status              developmentCardStatus
}

func (r DevelopmentCard) GetType() developmentCardType {
	return r.developmentCardType
}

func (r DevelopmentCard) GetStatus() developmentCardStatus {
	return r.status
}
