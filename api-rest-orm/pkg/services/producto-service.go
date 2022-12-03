package services

import (
	"fmt"
	apiRest "orm"
	"orm/pkg/logger"
	"os"
)

// Service Desacoplar la capa de servicio/
type Service interface {
	CreateProducto(producto *apiRest.Producto) error

	GetProductos() (*[]apiRest.Producto, error)

	DeleteProducto(ID uint) (error, int64)

	UpdateProducto(ID string, producto *apiRest.Producto) error

	GetProductoByID(ID uint) (*apiRest.Producto, error)
}

type service struct {
	repository apiRest.Repository
	log        logger.Logger
}

func NewService(repository apiRest.Repository) Service {
	return &service{repository: repository, log: logger.NewLogger(os.Getenv(apiRest.LOG_LEVEL))}
}

func (s *service) CreateProducto(producto *apiRest.Producto) error {
	return s.repository.CreateProducto(producto)
}

func (s *service) GetProductos() (*[]apiRest.Producto, error) {
	return s.repository.GetProducto()
}

func (s *service) DeleteProducto(ID uint) (error, int64) {
	s.log.Debug(fmt.Sprintf("Deleted product with ID [%v]", ID))
	return s.repository.DeleteProducto(ID)
}

func (s *service) UpdateProducto(ID string, producto *apiRest.Producto) error {
	return s.repository.UpdateProducto(ID, producto)
}

func (s *service) GetProductoByID(ID uint) (*apiRest.Producto, error) {
	return s.repository.GetProductoByID(ID)
}
