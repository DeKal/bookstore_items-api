package services

import (
	"net/http"

	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dao"
	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dto"
	"github.com/DeKal/bookstore_utils-go/errors"
)

// ItemsServiceInterface is an interface for items service
type ItemsServiceInterface interface {
	Create(*dto.Item) (*dto.Item, *errors.RestError)
	Get(string) (*dto.Item, *errors.RestError)
}

type itemsService struct {
	itemDAO dao.ItemsDAOInterface
}

// NewItemsService return new items service
func NewItemsService(itemDAO dao.ItemsDAOInterface) ItemsServiceInterface {
	return &itemsService{
		itemDAO: itemDAO,
	}
}

func (s *itemsService) Create(item *dto.Item) (*dto.Item, *errors.RestError) {
	return s.itemDAO.Save(item)
}

func (s *itemsService) Get(string) (*dto.Item, *errors.RestError) {
	return nil, &errors.RestError{
		Status:  http.StatusNotImplemented,
		Message: "Implement me",
		Error:   "not_implement",
	}
}
