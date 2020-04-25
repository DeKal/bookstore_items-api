package services

import (
	"net/http"

	"github.com/DeKal/bookstore_items-api/pkg/domain/items"
	"github.com/DeKal/bookstore_utils-go/errors"
)

// ItemsServiceInterface is an interface for items service
type ItemsServiceInterface interface {
	Create(*items.Item) (*items.Item, *errors.RestError)
	Get(string) (*items.Item, *errors.RestError)
}

type itemsService struct {
}

// NewItemsService return new items service
func NewItemsService() ItemsServiceInterface {
	return &itemsService{}
}

func (s *itemsService) Create(item *items.Item) (*items.Item, *errors.RestError) {
	return nil, &errors.RestError{
		Status:  http.StatusNotImplemented,
		Message: "Implement me",
		Error:   "not_implement",
	}
}

func (s *itemsService) Get(string) (*items.Item, *errors.RestError) {
	return nil, &errors.RestError{
		Status:  http.StatusNotImplemented,
		Message: "Implement me",
		Error:   "not_implement",
	}
}
