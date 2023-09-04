package processor

import (
	"math/rand"
	"reflect"
	"rps/model"
	"testing"

	"github.com/google/uuid"
)

func TestProcessor_CalculatePoints(t *testing.T) {
	uuid.SetRand(rand.New(rand.NewSource(1)))
	// this is a deterministic value because we are seeding the uuid generator.
	expectedID := "52fdfc07-2182-454f-963f-5f0f9a621d72"
	type args struct {
		request model.ReceiptRequest
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want model.ReceiptResponse
	}{
		{
			name: "Valid receipt returns response with ID",
			args: args{
				request: model.ReceiptRequest{
					Retailer:     "retailer",
					PurchaseDate: "2023-09-01",
					PurchaseTime: "14:55",
					Items: []model.Item{
						{
							ShortDescription: "some item",
							Price:            "3.99",
						},
					},
					Total: "45.00",
				},
			},
			want: model.ReceiptResponse{
				ID: expectedID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.CalculatePoints(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processor.CalculatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_GetPoints(t *testing.T) {
	type args struct {
		receiptID string
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want model.PointsResponse
	}{
		{
			name: "ID does not exist",
			args: args{
				receiptID: "some-id",
			},
			want: model.PointsResponse{Points: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.GetPoints(tt.args.receiptID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processor.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_calcPointsRetailerName(t *testing.T) {
	type args struct {
		retailer string
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want int
	}{
		{
			name: "Retailer name is alphanumeric",
			args: args{
				retailer: "retailer name length is 22",
			},
			want: 22,
		},
		{
			name: "Retailer name has special characters",
			args: args{
				retailer: "    re -+_!@#$%^&*() tailer   ",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.calcPointsRetailerName(tt.args.retailer); got != tt.want {
				t.Errorf("Processor.calcPointsRetailerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_calcPointsReceiptTotal(t *testing.T) {
	type args struct {
		total string
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want int
	}{
		{
			name: "total is a round dollar amount",
			args: args{
				total: "42.00",
			},
			want: 75,
		},
		{
			name: "total has a decimal amount that is a multiple of 0.25",
			args: args{
				total: "42.25",
			},
			want: 25,
		},
		{
			name: "total has a decimal amount that is not a multiple of 0.25",
			args: args{
				total: "42.03",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.calcPointsReceiptTotal(tt.args.total); got != tt.want {
				t.Errorf("Processor.calcPointsReceiptTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_calcPointsItems(t *testing.T) {
	type args struct {
		items []model.Item
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want int
	}{
		{
			name: "One item having length not a multiple of 3",
			args: args{
				items: []model.Item{
					{
						ShortDescription: "Gatorade",
						Price:            "4.00",
					},
				},
			},
			want: 0,
		},
		{
			name: "One item having length of 6 characters",
			args: args{
				items: []model.Item{
					{
						ShortDescription: "Orange",
						Price:            "1.99",
					},
				},
			},
			want: 1,
		},
		{
			name: "One item having length as 6 characters and extra spaces",
			args: args{
				items: []model.Item{
					{
						ShortDescription: " Red Apple ",
						Price:            "1.99",
					},
				},
			},
			want: 1,
		},
		{
			name: "Two items",
			args: args{
				items: []model.Item{
					{
						ShortDescription: "First",
						Price:            "3.00",
					},
					{
						ShortDescription: "Second Item",
						Price:            "7.49",
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.calcPointsItems(tt.args.items); got != tt.want {
				t.Errorf("Processor.calcPointsItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_calcPointsDateTime(t *testing.T) {
	type args struct {
		purchaseDate string
		purchaseTime string
	}
	tests := []struct {
		name string
		p    *Processor
		args args
		want int
	}{
		{
			name: "Odd day, and hour after 2p and before 4p",
			args: args{
				purchaseDate: "2023-09-01",
				purchaseTime: "14:03",
			},
			want: 16,
		},
		{
			name: "Odd day, and hour in the morning",
			args: args{
				purchaseDate: "2023-09-01",
				purchaseTime: "11:11",
			},
			want: 6,
		},
		{
			name: "Even day, and hour after 2pm and before 4pm",
			args: args{
				purchaseDate: "2023-09-02",
				purchaseTime: "14:11",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.calcPointsDateTime(tt.args.purchaseDate, tt.args.purchaseTime); got != tt.want {
				t.Errorf("Processor.calcPointsDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
