package client

import (
	"reflect"
	"testing"

	"example.com/m/external/dto"
)

func TestNewProductClient(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want *ProductClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductClient(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductClient_FetchProductDetails(t *testing.T) {
	type args struct {
		productID int
	}
	tests := []struct {
		name    string
		pc      *ProductClient
		args    args
		want    *dto.ProductDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pc.FetchProductDetails(tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductClient.FetchProductDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductClient.FetchProductDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
