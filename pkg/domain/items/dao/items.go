package dao

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DeKal/bookstore_items-api/pkg/clients/elasticsearch"
	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dto"
	"github.com/DeKal/bookstore_utils-go/errors"
)

const (
	index    = "items"
	typeItem = "item"
)

type itemsDAO struct {
	esClient elasticsearch.ClientInterface
}

// ItemsDAOInterface provides methods for interacting with DB
type ItemsDAOInterface interface {
	Save(*dto.Item) (*dto.Item, *errors.RestError)
	Get(string) (*dto.Item, *errors.RestError)
}

// NewItemsDao return new NewItemsDao interface
func NewItemsDao(esClient elasticsearch.ClientInterface) ItemsDAOInterface {
	return &itemsDAO{
		esClient: esClient,
	}
}

func (dao *itemsDAO) Save(item *dto.Item) (*dto.Item, *errors.RestError) {
	saveResult, err := dao.esClient.IndexWith(index, typeItem, item)
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to save item.")
	}
	item.ID = saveResult.Id
	return item, nil
}

func (dao *itemsDAO) Get(id string) (*dto.Item, *errors.RestError) {
	doc, err := dao.esClient.Get(index, typeItem, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return nil, errors.NewNotFoundError(fmt.Sprintf("no items found with id %s", id))
		}
		return nil, errors.NewInternalServerError(fmt.Sprintf("Error when trying to get id %s", id))
	}

	bytes, err := doc.Source.MarshalJSON()
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to parse db response.")
	}

	item := &dto.Item{}
	if err := json.Unmarshal(bytes, item); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to parse db response.")
	}
	return item, nil
}
