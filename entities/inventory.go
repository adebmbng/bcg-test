package entities

import "time"

type Inventory struct {
	Id        int        `json:"id"`
	Sku       string     `json:"sku"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	Qty       int        `json:"qty"`
	IsActive  int        `json:"isActive"`
	CreatedAt *time.Time `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy string     `json:"updatedBy"`
}

type GetPriceResponse struct {
	TotalPrice          float32 `json:"totalPrice"`
	DiscountPrice       float32 `json:"discountPrice"`
	FreeItem            string  `json:"freeItem"`
	FinalPrice          float32 `json:"finalPrice"`
	FormattedFinalPrice string  `json:"formattedFinalPrice"`
}

type GetPriceRequest struct {
	Items map[string]int
}

// NewGetPriceRequest to convert from string request to price request
func NewGetPriceRequest(q []string) *GetPriceRequest {
	ret := make(map[string]int)
	for _, s := range q {
		ret[s]++
	}
	return &GetPriceRequest{
		Items: ret,
	}
}

// GetSKUs to get list of sku
func (g *GetPriceRequest) GetSKUs() []string {
	skus := make([]string, 0)
	for k := range g.Items {
		skus = append(skus, k)
	}
	return skus
}
