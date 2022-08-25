package models

type Construction struct {
	aggregate
	constructionType constructionType
	land             *Land
}

func (c Construction) GetType() constructionType {
	return c.constructionType
}

func (c Construction) GetLand() *Land {
	return c.land
}
