package elasticsearch

import (
	"fmt"
	"time"

	"github.com/DeKal/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

// Client contains elasticsearch client
type Client struct {
	client *elastic.Client
}

// ClientInterface define interface for getting *elastic.Client
type ClientInterface interface {
	IndexWith(string, interface{}) (*elastic.IndexResponse, error)
}

// NewEsClient return new elasticsearch client
func NewEsClient() ClientInterface {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
	)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: client,
	}
}

// IndexWith index a document to es with specified index and doc
func (c *Client) IndexWith(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).Type(index).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}
