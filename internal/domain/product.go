package domain

type Product struct {
	Id_product  int    `json:"id_product"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type Products []Product
