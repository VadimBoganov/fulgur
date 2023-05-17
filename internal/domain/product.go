package domain

type Product struct {
	Id   int
	Name string
}

type ProductType struct {
	Id      int
	Product Product
	Name    string
}

type ProductSubType struct {
	Id          int
	ProductType ProductType
	Name        string
}
