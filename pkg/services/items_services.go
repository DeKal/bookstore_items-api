package services

import (
	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dao"
	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dto"
	"github.com/DeKal/bookstore_items-api/pkg/domain/queries"
	"github.com/DeKal/bookstore_utils-go/errors"
)

// ItemsServiceInterface is an interface for items service
type ItemsServiceInterface interface {
	Create(*dto.Item) (*dto.Item, *errors.RestError)
	Get(string) (*dto.Item, *errors.RestError)
	Seach(query *queries.EsQuery) ([]dto.Item, *errors.RestError)
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

func (s *itemsService) Get(id string) (*dto.Item, *errors.RestError) {
	return s.itemDAO.Get(id)
}

func (s *itemsService) Seach(query *queries.EsQuery) ([]dto.Item, *errors.RestError) {
	return s.itemDAO.Search(query)
}
