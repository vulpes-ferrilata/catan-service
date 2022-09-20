package models

type Construction struct {
	aggregate
	constructionType ConstructionType
	land             *Land
}

func (c Construction) GetType() ConstructionType {
	return c.constructionType
}

func (c Construction) GetLand() *Land {
	return c.land
}
