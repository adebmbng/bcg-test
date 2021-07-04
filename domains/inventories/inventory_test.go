package inventories

import (
	"github.com/adebmbng/bcg-test/entities"
	"reflect"
	"testing"
)

func Test_validateInventoriesBySKU(t *testing.T) {
	mapData := map[string]*entities.Inventory{
		"1": {
			Id:       1,
			Sku:      "1",
			Name:     "1",
			Price:    1,
			Qty:      1,
			IsActive: 1,
		},
		"2": {
			Id:       2,
			Sku:      "2",
			Name:     "2",
			Price:    2,
			Qty:      2,
			IsActive: 1,
		},
	}
	type args struct {
		inventoryMap map[string]*entities.Inventory
		q            []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "#1 - positive",
			args: args{
				inventoryMap: mapData,
				q:            []string{"1", "2"},
			},
			want: true,
		},
		{
			name: "#2 - negative",
			args: args{
				inventoryMap: mapData,
				q:            []string{"1", "2", "3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateInventoriesBySKU(tt.args.inventoryMap, tt.args.q); got != tt.want {
				t.Errorf("validateInventoriesBySKU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFinalPrice(t *testing.T) {
	macbookPro := &entities.Inventory{
		Id:       1,
		Sku:      "43N23P",
		Name:     "Macbook Pro",
		Price:    5399.99,
		Qty:      10,
		IsActive: 1,
	}
	googleHome := &entities.Inventory{
		Id:       2,
		Sku:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Qty:      10,
		IsActive: 1,
	}
	alexaSpeaker := &entities.Inventory{
		Id:       3,
		Sku:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.5,
		Qty:      10,
		IsActive: 1,
	}
	raspi := &entities.Inventory{
		Id:       4,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30,
		Qty:      2,
		IsActive: 1,
	}

	freeItem := &entities.Promotion{
		Id:            1,
		Item:          "43N23P",
		PromotionType: entities.PromotionType__FreeItem,
		MinimumQty:    1,
		PromotionData: "234234",
		IsActive:      1,
	}
	buy3pay2 := &entities.Promotion{
		Id:            2,
		Item:          "120P90",
		PromotionType: entities.PromotionType__Buy3Pay2,
		MinimumQty:    3,
		PromotionData: "",
		IsActive:      1,
	}
	discount := &entities.Promotion{
		Id:            3,
		Item:          "A304SD",
		PromotionType: entities.PromotionType__Discount,
		MinimumQty:    3,
		PromotionData: "10",
		IsActive:      1,
	}
	promos := map[string][]*entities.Promotion{
		freeItem.Item: {freeItem},
		buy3pay2.Item: {buy3pay2},
		discount.Item: {discount},
	}
	inventories := map[string]*entities.Inventory{
		macbookPro.Sku:   macbookPro,
		raspi.Sku:        raspi,
		googleHome.Sku:   googleHome,
		alexaSpeaker.Sku: alexaSpeaker,
	}
	type args struct {
		req          *entities.GetPriceRequest
		prms         map[string][]*entities.Promotion
		inventoryMap map[string]*entities.Inventory
	}
	tests := []struct {
		name string
		args args
		want entities.GetPriceResponse
	}{
		{
			name: "mbp, raspi",
			args: args{
				req: &entities.GetPriceRequest{map[string]int{
					macbookPro.Sku: 1,
					raspi.Sku:      1,
				}},
				prms:         promos,
				inventoryMap: inventories,
			},
			want: entities.GetPriceResponse{
				TotalPrice:          macbookPro.Price + raspi.Price,
				DiscountPrice:       raspi.Price,
				FreeItem:            raspi.Sku,
				FinalPrice:          macbookPro.Price,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "gh, gh, gh",
			args: args{
				req: &entities.GetPriceRequest{map[string]int{
					googleHome.Sku: 3,
				}},
				prms:         promos,
				inventoryMap: inventories,
			},
			want: entities.GetPriceResponse{
				TotalPrice:          googleHome.Price * 3,
				DiscountPrice:       googleHome.Price,
				FinalPrice:          googleHome.Price * 2,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "alexa, alexa, alexa",
			args: args{
				req: &entities.GetPriceRequest{map[string]int{
					alexaSpeaker.Sku: 3,
				}},
				prms:         promos,
				inventoryMap: inventories,
			},
			want: entities.GetPriceResponse{
				TotalPrice:          alexaSpeaker.Price * 3,
				DiscountPrice:       alexaSpeaker.Price * 0.1 * 3,
				FinalPrice:          (alexaSpeaker.Price * 3) - alexaSpeaker.Price*0.1*3,
				FormattedFinalPrice: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFinalPrice(tt.args.req, tt.args.prms, tt.args.inventoryMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateFinalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
