package elasticsearch

import (
	"fmt"
	"time"

	"github.com/DeKal/bookstore_utils-go/logger"
	"github.com/DeKal/bookstore_utils-go/logger/elasticlogger"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

// Client contains elasticsearch client
type Client struct {
	client *elastic.Client
}

// ClientInterface define interface for getting *elastic.Client
type ClientInterface interface {
	IndexWith(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
}

// NewEsClient return new elasticsearch client
func NewEsClient() ClientInterface {

	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetInfoLog(elasticlogger.NewLogger(elasticlogger.LoggerInfo)),
		elastic.SetErrorLog(elasticlogger.NewLogger(elasticlogger.LoggerError)),
	)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: client,
	}
}

// IndexWith index a document to es with specified index and doc
func (c *Client) IndexWith(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

// Get return document from es
func (c *Client) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(
			fmt.Sprintf("Error when trying to get document in index %s with type %s", index, docType), err)
		return nil, err
	}
	return result, nil
}
