package common

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerModule struct {
	Handler      func(*fiber.Ctx) error
	Route        string
	Method       interface{}
	RequiresAuth bool
}

type SliceHandlers struct {
	Prefix string
	Routes []HandlerModule
}
type GlobalHandlers []SliceHandlers

type HandlersStore struct {
	Handlers []SliceHandlers
}

func NewHandlersStore() *HandlersStore {
	return &HandlersStore{}
}
