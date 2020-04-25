package controllers

import (
	"fmt"
	"net/http"

	"github.com/DeKal/bookstore_items-api/pkg/domain/items"
	"github.com/DeKal/bookstore_items-api/pkg/services"
	"github.com/DeKal/bookstore_oauth-go/oauth"
)

// ItemsControllerInterface is an interface for itemService
type ItemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
	services services.ItemsServiceInterface
}

// NewItemsController return new itemService
func NewItemsController() ItemsControllerInterface {
	return &itemsController{}
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := c.services.Create(item)
	if err != nil {
		return
	}

	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}
}
