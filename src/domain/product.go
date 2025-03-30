package domain

import "time"

// Product representa un producto en MongoDB
type Product struct {
	SKU                 string    `bson:"sku" json:"sku"`
	Nombre              string    `bson:"nombre" json:"nombre"`
	Cantidad_disponible float64   `bson:"cantidad_disponible" json:"cantidad_disponible"`
	Fecha_entrada       time.Time `bson:"fecha_entrada" json:"fecha_entrada"`
}

// ProductRepository define el contrato para almacenar productos
type ProductRepository interface {
	SaveProducts(products []Product) error
	GetProducts() ([]Product, error)
}
