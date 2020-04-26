package dao

import (
	"github.com/DeKal/bookstore_items-api/pkg/clients/elasticsearch"
	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dto"
	"github.com/DeKal/bookstore_utils-go/errors"
)

type itemsDAO struct {
	esClient elasticsearch.ClientInterface
}

// ItemsDAOInterface provides methods for interacting with DB
type ItemsDAOInterface interface {
	Save(*dto.Item) (*dto.Item, *errors.RestError)
}

// NewItemsDao return new NewItemsDao interface
func NewItemsDao(esClient elasticsearch.ClientInterface) ItemsDAOInterface {
	return &itemsDAO{
		esClient: esClient,
	}
}

func (dao *itemsDAO) Save(item *dto.Item) (*dto.Item, *errors.RestError) {
	saveResult, err := dao.esClient.IndexWith("items", item)
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to save item.")
	}
	item.ID = saveResult.Id
	return item, nil
}
