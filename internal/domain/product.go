package domain

// металы, припои
type Product struct {
	Id   int
	Name string
}

// Оловянно-свинцовые припои, Индиевые припои
type ProductType struct {
	Id        int
	ProductId int
	Name      string
}

// Малосурьмянистые, Сурьмянистые
type ProductSubType struct {
	Id          int
	ProductType ProductType
	Name        string
}
