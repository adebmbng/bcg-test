package promos

import (
	"github.com/adebmbng/bcg-test/entities"
	errwrapper "github.com/adebmbng/bcg-test/pkg/error"
	"github.com/adebmbng/bcg-test/repositories/mysql"
)

type promo struct {
	repo mysql.Repository
}

type Promo interface {
	GetAvailablePromoBySKUs(q []string) (map[string][]*entities.Promotion, error)
}

func NewPromos(repo mysql.Repository) Promo {
	return &promo{
		repo: repo,
	}
}

// GetAvailablePromoBySKUs to get list of promo of each SKU
func (p *promo) GetAvailablePromoBySKUs(q []string) (map[string][]*entities.Promotion, error) {
	// validate request
	if len(q) == 0 {
		err := errwrapper.ErrorBadRequest
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.BadRequest)
	}

	// get promo from db
	promos, err := p.repo.GetPromoBySKUs(q)
	if err != nil {
		return nil, errwrapper.ErrWrap(err, err.Error(), errwrapper.RepositoryError)
	}

	// map promos to return param
	ret := make(map[string][]*entities.Promotion)
	// iterate promos
	for _, promotion := range promos {
		ret[promotion.Item] = append(ret[promotion.Item], promotion)
	}

	return ret, nil
}
