package inventories

import (
	"context"
	"github.com/adebmbng/bcg-test/domains/promos"
	"github.com/adebmbng/bcg-test/entities"
	errwrapper "github.com/adebmbng/bcg-test/pkg/error"
	"github.com/adebmbng/bcg-test/repositories/mysql"
	"github.com/leekchan/accounting"
)

type inventory struct {
	repo        mysql.Repository
	promoDomain promos.Promo
}

type Inventory interface {
	GetPrice(ctx context.Context, q []string) (*entities.GetPriceResponse, error)
}

func NewInventories(repo mysql.Repository, promoDomain promos.Promo) Inventory {
	return &inventory{
		repo:        repo,
		promoDomain: promoDomain,
	}
}

// GetPrice to get price by scanned items (string)
func (i *inventory) GetPrice(ctx context.Context, q []string) (*entities.GetPriceResponse, error) {
	// validate request
	if len(q) == 0 {
		err := errwrapper.ErrorBadRequest
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.BadRequest)
	}

	// convert request to GetPriceRequest
	req := entities.NewGetPriceRequest(q)

	// get every detail of the products
	inventories, err := i.repo.GetInventoriesBySKUs(req.GetSKUs())
	if err != nil {
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.RepositoryError)
	}
	// create map to validate
	inventoryMap := make(map[string]*entities.Inventory)
	// iterate inventories
	for _, in := range inventories {
		inventoryMap[in.Sku] = in
	}

	// validate if all the request sku is found
	if validateInventoriesBySKU(inventoryMap, req.GetSKUs()) {
		err := errwrapper.ErrorBadRequest
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.BadRequest)
	}

	// get available promos
	prms, err := i.promoDomain.GetAvailablePromoBySKUs(req.GetSKUs())
	if err != nil {
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.RepositoryError)
	}

	ret := calculateFinalPrice(req, prms, inventoryMap)

	// format total price
	ac := accounting.Accounting{Symbol: "$ ", Precision: 2}
	ret.FormattedFinalPrice = ac.FormatMoney(ret.TotalPrice)

	return &ret, nil
}

// calculateFinalPrice to do calculating the final price
func calculateFinalPrice(req *entities.GetPriceRequest, prms map[string][]*entities.Promotion, inventoryMap map[string]*entities.Inventory) entities.GetPriceResponse {
	// create default variable for return value
	var ret entities.GetPriceResponse
	// iterate req to sum the prize
	for k, v := range req.Items {
		// validate if it has promo
		ps := prms[k]
		// TODO: Promo should have priority rank
		// iterate promos to get eligible promo
		var p *entities.GetPriceResponse
		for _, promotion := range ps {
			p = promotion.IsEligible(inventoryMap[k], v)
		}
		if p != nil {
			ret.TotalPrice += p.TotalPrice
			ret.DiscountPrice += p.DiscountPrice
			ret.FinalPrice += p.FinalPrice
			// TODO: decide what free item if has multiple free item
			if p.FreeItem != `` {
				ret.FreeItem = p.FreeItem
			}
		} else {
			ret.TotalPrice += inventoryMap[k].Price * float32(v)
			ret.FinalPrice += inventoryMap[k].Price * float32(v)
		}
	}

	// make the free item price become 0
	if ret.FreeItem != `` {
		ret.FinalPrice -= inventoryMap[ret.FreeItem].Price
		ret.DiscountPrice += inventoryMap[ret.FreeItem].Price
	}

	return ret
}

// validateInventoriesBySKU to validate each request has inventory data
func validateInventoriesBySKU(inventoryMap map[string]*entities.Inventory, q []string) bool {

	// validate each map by the request
	for _, s := range q {
		// if inventory is null it means that something is mission
		if inventoryMap[s] == nil {
			return false
		}
	}

	// everything is fine
	return true
}
