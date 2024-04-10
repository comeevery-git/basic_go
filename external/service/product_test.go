package service

import (
	"reflect"
	"testing"

	"example.com/m/external/client"
	"example.com/m/external/dto"
)

func TestNewProductService(t *testing.T) {
	type args struct {
		productClient *client.ProductClient
	}
	tests := []struct {
		name string
		args args
		want ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.productClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productService_GetProductDetails(t *testing.T) {
	type args struct {
		productID int
	}
	tests := []struct {
		name    string
		s       *productService
		args    args
		want    *dto.ProductDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetProductDetails(tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("productService.GetProductDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productService.GetProductDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
