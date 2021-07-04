package entities

import (
	"strconv"
	"time"
)

type PromotionType string

const (
	PromotionType__FreeItem PromotionType = `free-item`
	PromotionType__Buy3Pay2 PromotionType = `buy-3-pay-2`
	PromotionType__Discount PromotionType = `discount`

	Buy3        int     = 3
	MakePercent float32 = 100
)

type Promotion struct {
	Id            int           `json:"id"`
	Item          string        `json:"item"`
	PromotionType PromotionType `json:"promotionType"`
	MinimumQty    int           `json:"minimumQty"`
	PromotionData string        `json:"promotionData"`
	IsActive      int           `json:"isActive"`
	CreatedAt     *time.Time    `json:"createdAt"`
	CreatedBy     string        `json:"createdBy"`
	UpdatedAt     *time.Time    `json:"updatedAt"`
	UpdatedBy     string        `json:"updatedBy"`
}

// IsEligible is to validate promo eligible
func (p *Promotion) IsEligible(item *Inventory, qty int) *GetPriceResponse {
	// validate sku
	if p.Item != item.Sku {
		return nil
	}
	// validate minimum qty
	if p.MinimumQty > qty {
		return nil
	}
	switch p.PromotionType {
	case PromotionType__FreeItem:
		// return as free item
		return &GetPriceResponse{
			TotalPrice:    item.Price * float32(qty),
			DiscountPrice: 0,
			FreeItem:      p.PromotionData,
			FinalPrice:    item.Price * float32(qty),
		}
	case PromotionType__Buy3Pay2:
		// use div to get free discount
		qtyFree := qty / Buy3

		// return as buy 3 pay 2
		return &GetPriceResponse{
			TotalPrice:    item.Price * float32(qty),
			DiscountPrice: item.Price * float32(qtyFree),
			FinalPrice:    item.Price * float32(qty-qtyFree),
		}
	case PromotionType__Discount:
		discountValue, err := strconv.Atoi(p.PromotionData)
		if err != nil {
			return nil
		}
		// return as discount price
		return &GetPriceResponse{
			TotalPrice:    item.Price * float32(qty),
			DiscountPrice: item.Price * (float32(discountValue) / MakePercent) * float32(qty),
			FinalPrice:    (item.Price * float32(qty)) - (item.Price * (float32(discountValue) / MakePercent) * float32(qty)),
		}
	}
	return nil
}
