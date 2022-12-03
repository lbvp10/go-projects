package postgresql

import (
	"gorm.io/gorm"
	apiRest "orm"
)

type postgresql struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) apiRest.Repository {
	return &postgresql{db: db}
}

func (p *postgresql) CreateProducto(producto *apiRest.Producto) error {
	result := p.db.Create(&producto)
	return result.Error
}

func (p *postgresql) GetProducto() (*[]apiRest.Producto, error) {
	var productos []apiRest.Producto
	result := p.db.Find(&productos)

	return &productos, result.Error
}

func (p *postgresql) DeleteProducto(ID uint) (error, int64) {
	result := p.db.Delete(&apiRest.Producto{}, ID)
	return result.Error, result.RowsAffected
}

func (p *postgresql) UpdateProducto(ID string, producto *apiRest.Producto) error {
	result := p.db.Save(&producto)
	return result.Error
}

func (p *postgresql) GetProductoByID(ID uint) (*apiRest.Producto, error) {
	producto := apiRest.Producto{Model: gorm.Model{ID: ID}}
	result := p.db.First(&producto)

	return &producto, result.Error
}
