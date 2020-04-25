package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DeKal/bookstore_items-api/pkg/domain/items"
	"github.com/DeKal/bookstore_items-api/pkg/services"
	"github.com/DeKal/bookstore_items-api/pkg/utils/httputils"
	"github.com/DeKal/bookstore_oauth-go/oauth"
	"github.com/DeKal/bookstore_utils-go/errors"
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
		httputils.WriteReponseError(w, err)
		return
	}

	item := &items.Item{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.WriteReponseError(w,
			errors.NewBadRequestError("Cannot parse request body for request item."))
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(requestBody, item); err != nil {
		httputils.WriteReponseError(w,
			errors.NewBadRequestError("Cannot parse request body to json."))
		return
	}
	item.Seller = oauth.GetCallerID(r)

	result, svcErr := c.services.Create(item)
	if svcErr != nil {
		httputils.WriteReponseError(w, svcErr)
		return
	}

	httputils.WriteJSONResponse(w, http.StatusCreated, result)
	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}
}
