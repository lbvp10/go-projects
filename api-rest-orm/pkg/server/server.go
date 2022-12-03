package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	api "orm"
	"orm/pkg/logger"
	"orm/pkg/services"
)

type Server interface {
	configServer()
}

func NewServer(ID string, log logger.Logger, app *fiber.App, productoService services.Service) Server {
	s := &server{
		serverID:        ID,
		productoService: productoService,
		app:             app,
		logger:          log,
	}
	s.configServer()
	return s
}

type server struct {
	app             *fiber.App
	productoService services.Service
	serverID        string
	logger          logger.Logger
}

func (s *server) configServer() {
	logsMiddleware := NewLogsMiddleware(s.logger)
	// Router
	api := s.app.Group("/api", requestid.New(), logsMiddleware.LoggerApiRequest)

	api.Get("/", s.getHandler)
	api.Get("/:id<int>", s.getByIdHandler)
	api.Post("/", s.postHandler)
	api.Put("/", s.putHandler)
	api.Delete("/:id<int>", s.deleteHandler)
}

func (s *server) getByIdHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	producto, err := s.productoService.GetProductoByID(uint(id))
	if err != nil {
		errorApi := logger.NewInternalServerApiError(err, "Error getById de la postgresql", s.logger)
		return ctx.Status(errorApi.Status).JSON(errorApi)
	} else if producto == nil {
		errorApi := logger.NewNotFoundApiError("")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}

func (s *server) getHandler(ctx *fiber.Ctx) error {
	productos, err := s.productoService.GetProductos()
	if err != nil {
		errorApi := logger.NewInternalServerApiError(err, "api consultando los productos", s.logger)
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(productos)

}
func (s *server) postHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	err := s.productoService.CreateProducto(producto)
	if err != nil {
		errorApi := logger.NewInternalServerApiError(err, "Error guardando en la postgresql", s.logger)
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}
func (s *server) putHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	err := s.productoService.UpdateProducto(string(1), producto)
	if err.Error != nil {
		errorApi := logger.NewInternalServerApiError(err, "Error actualizando la postgresql", s.logger)
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}

func (s *server) deleteHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err, rowsDeleted := s.productoService.DeleteProducto(uint(id))

	if err != nil {
		errorApi := logger.NewInternalServerApiError(err, "Error Delete", s.logger)
		return ctx.Status(errorApi.Status).JSON(errorApi)
	} else if rowsDeleted < 1 {
		errorApi := logger.NewNotFoundApiError("Delete not found")
		return ctx.Status(errorApi.Status).JSON(errorApi)

	}
	return ctx.SendStatus(200)

}

func loadProductByRequest(c *fiber.Ctx) (*api.Producto, error) {
	producto := api.Producto{}
	var err error
	if err = c.BodyParser(&producto); err != nil {
		errorApi := logger.NewBadRequestApiError("Error body")
		return nil, c.Status(errorApi.Status).JSON(errorApi)
	}

	return &producto, err
}
