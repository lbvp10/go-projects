package api_rest_orm

import (
	"gorm.io/gorm"
	"time"
)

type Producto struct {
	gorm.Model
	Fecha       time.Time `json:"fecha"`
	Descripcion string    `json:"descripcion"`
}

// Repository Desacoplar la capa de persistencia/
type Repository interface {
	CreateProducto(producto *Producto) error

	GetProducto() (*[]Producto, error)

	DeleteProducto(ID uint) (error, int64)

	UpdateProducto(ID string, producto *Producto) error

	GetProductoByID(ID uint) (*Producto, error)
}
