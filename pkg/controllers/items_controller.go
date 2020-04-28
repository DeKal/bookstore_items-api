package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/DeKal/bookstore_items-api/pkg/domain/items/dto"
	"github.com/DeKal/bookstore_items-api/pkg/domain/queries"
	"github.com/DeKal/bookstore_items-api/pkg/services"
	"github.com/DeKal/bookstore_items-api/pkg/utils/httputils"
	"github.com/DeKal/bookstore_oauth-go/oauth"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/gorilla/mux"
)

// ItemsControllerInterface is an interface for itemService
type ItemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct {
	services services.ItemsServiceInterface
}

// NewItemsController return new itemService
func NewItemsController(services services.ItemsServiceInterface) ItemsControllerInterface {
	return &itemsController{
		services: services,
	}
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		httputils.WriteReponseError(w, err)
		return
	}
	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		httputils.WriteReponseError(w,
			errors.NewUnauthorizedError("Unauthorized with given access_token"))
		return
	}

	item := &dto.Item{}
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
	item.Seller = sellerID

	result, svcErr := c.services.Create(item)
	if svcErr != nil {
		httputils.WriteReponseError(w, svcErr)
		return
	}

	httputils.WriteJSONResponse(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	if len(itemID) == 0 {
		httputils.WriteReponseError(w,
			errors.NewBadRequestError("No item ID from request."))
		return
	}

	item, err := c.services.Get(itemID)
	if err != nil {
		httputils.WriteReponseError(w, err)
		return
	}

	httputils.WriteJSONResponse(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.WriteReponseError(w,
			errors.NewBadRequestError("Cannot parse search query."))
		return
	}
	defer r.Body.Close()

	query := &queries.EsQuery{}
	if err := json.Unmarshal(bytes, query); err != nil {
		httputils.WriteReponseError(w,
			errors.NewBadRequestError("Cannot parse search query."))
		return
	}

	items, searchErr := c.services.Seach(query)
	if searchErr != nil {
		httputils.WriteReponseError(w, searchErr)
		return
	}

	httputils.WriteJSONResponse(w, http.StatusOK, items)
}
