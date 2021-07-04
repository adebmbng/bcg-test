package entities

import (
	"reflect"
	"testing"
	"time"
)

func TestPromotion_IsEligible(t *testing.T) {
	type fields struct {
		Id            int
		Item          string
		PromotionType PromotionType
		MinimumQty    int
		PromotionData string
		IsActive      int
		CreatedAt     *time.Time
		CreatedBy     string
		UpdatedAt     *time.Time
		UpdatedBy     string
	}
	type args struct {
		item *Inventory
		qty  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetPriceResponse
	}{
		{
			name: "buy macbook free raspi - positive",
			fields: fields{
				Id:            0,
				Item:          "43N23P",
				PromotionType: PromotionType__FreeItem,
				MinimumQty:    1,
				PromotionData: "234234",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "43N23P",
					Name:     "Macbook Pro",
					Price:    5399.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 1,
			},
			want: &GetPriceResponse{
				TotalPrice:          5399.99,
				DiscountPrice:       0,
				FreeItem:            "234234",
				FinalPrice:          5399.99,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy macbook free raspi - qty not meet the requirement",
			fields: fields{
				Id:            0,
				Item:          "43N23P",
				PromotionType: PromotionType__FreeItem,
				MinimumQty:    2,
				PromotionData: "234234",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "43N23P",
					Name:     "Macbook Pro",
					Price:    5399.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 1,
			},
			want: nil,
		},
		{
			name: "buy macbook free raspi - not buy macbook",
			fields: fields{
				Id:            0,
				Item:          "43N23P",
				PromotionType: PromotionType__FreeItem,
				MinimumQty:    2,
				PromotionData: "234234",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "43N23P1",
					Name:     "Macbook Air",
					Price:    5399.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 1,
			},
			want: nil,
		},
		{
			name: "buy 3 google home pay 2 - positive",
			fields: fields{
				Id:            0,
				Item:          "120P90",
				PromotionType: PromotionType__Buy3Pay2,
				MinimumQty:    3,
				PromotionData: "",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 3,
			},
			want: &GetPriceResponse{
				TotalPrice:          49.99 * 3,
				DiscountPrice:       49.99,
				FreeItem:            "",
				FinalPrice:          49.99 * 2,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy 3 google home pay 2 - minimum qty not meet",
			fields: fields{
				Id:            0,
				Item:          "120P90",
				PromotionType: PromotionType__Buy3Pay2,
				MinimumQty:    3,
				PromotionData: "",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 2,
			},
			want: nil,
		},
		{
			name: "buy 3 google home pay 2 - buy 6 and should pay 4",
			fields: fields{
				Id:            0,
				Item:          "120P90",
				PromotionType: PromotionType__Buy3Pay2,
				MinimumQty:    3,
				PromotionData: "",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 6,
			},
			want: &GetPriceResponse{
				TotalPrice:          49.99 * 6,
				DiscountPrice:       49.99 * 2,
				FreeItem:            "",
				FinalPrice:          49.99 * 4,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy 3 google home pay 2 - buy 8 and should pay 4",
			fields: fields{
				Id:            0,
				Item:          "120P90",
				PromotionType: PromotionType__Buy3Pay2,
				MinimumQty:    3,
				PromotionData: "",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:       0,
					Sku:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Qty:      10,
					IsActive: 1,
				},
				qty: 8,
			},
			want: &GetPriceResponse{
				TotalPrice:          49.99 * 8,
				DiscountPrice:       49.99 * 2,
				FreeItem:            "",
				FinalPrice:          49.99 * 6,
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy 3 alexa get 10% discount - positive",
			fields: fields{
				Id:            0,
				Item:          "A304SD",
				PromotionType: PromotionType__Discount,
				MinimumQty:    3,
				PromotionData: "10",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:        0,
					Sku:       "A304SD",
					Name:      "Alexa Speaker",
					Price:     109.5,
					Qty:       10,
					IsActive:  1,
				},
				qty:  3,
			},
			want: &GetPriceResponse{
				TotalPrice:          109.5 * 3,
				DiscountPrice:       109.5 * 3 * 0.1,
				FreeItem:            "",
				FinalPrice:          109.5*3 - (109.5 * 3 * 0.1),
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy 3 alexa get 10% discount - buy 4",
			fields: fields{
				Id:            0,
				Item:          "A304SD",
				PromotionType: PromotionType__Discount,
				MinimumQty:    3,
				PromotionData: "10",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:        0,
					Sku:       "A304SD",
					Name:      "Alexa Speaker",
					Price:     109.5,
					Qty:       10,
					IsActive:  1,
				},
				qty:  4,
			},
			want: &GetPriceResponse{
				TotalPrice:          109.5 * 4,
				DiscountPrice:       109.5 * 4 * 0.1,
				FreeItem:            "",
				FinalPrice:          109.5*4 - (109.5 * 4 * 0.1),
				FormattedFinalPrice: "",
			},
		},
		{
			name: "buy 3 alexa get 10% discount - buy 2",
			fields: fields{
				Id:            0,
				Item:          "A304SD",
				PromotionType: PromotionType__Discount,
				MinimumQty:    3,
				PromotionData: "10",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:        0,
					Sku:       "A304SD",
					Name:      "Alexa Speaker",
					Price:     109.5,
					Qty:       10,
					IsActive:  1,
				},
				qty:  2,
			},
			want: nil,
		},
		{
			name: "buy 3 alexa get 10% discount - bad data",
			fields: fields{
				Id:            0,
				Item:          "A304SD",
				PromotionType: PromotionType__Discount,
				MinimumQty:    3,
				PromotionData: "10s",
				IsActive:      1,
			},
			args: args{
				item: &Inventory{
					Id:        0,
					Sku:       "A304SD",
					Name:      "Alexa Speaker",
					Price:     109.5,
					Qty:       10,
					IsActive:  1,
				},
				qty:  3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Promotion{
				Id:            tt.fields.Id,
				Item:          tt.fields.Item,
				PromotionType: tt.fields.PromotionType,
				MinimumQty:    tt.fields.MinimumQty,
				PromotionData: tt.fields.PromotionData,
				IsActive:      tt.fields.IsActive,
				CreatedAt:     tt.fields.CreatedAt,
				CreatedBy:     tt.fields.CreatedBy,
				UpdatedAt:     tt.fields.UpdatedAt,
				UpdatedBy:     tt.fields.UpdatedBy,
			}
			if got := p.IsEligible(tt.args.item, tt.args.qty); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEligible() = %v, want %v", got, tt.want)
			}
		})
	}
}
