package domain

// чушка, прошок
type Item struct {
	Id            int32
	ProductItemId int32
	Name          string
	Price         float64
	ImageUrl      string
	IsFullPrice   bool
}
