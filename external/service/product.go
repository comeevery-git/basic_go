package service

// 외부 서비스(예: 외부 API)와의 통신 로직 구현

import (
	"example.com/m/external/client"
	"example.com/m/external/dto"
)
type ProductService interface {
    GetProductDetails(productID int) (*dto.ProductDetails, error)
}

// ProductService 인터페이스의 구현체
type productService struct {
    productClient *client.ProductClient
}

// productService 인스턴스를 생성 및 반환
func NewProductService(productClient *client.ProductClient) ProductService {
    return &productService{productClient: productClient}
}

// 상품 정보 조회 by productID
func (s *productService) GetProductDetails(productID int) (*dto.ProductDetails, error) {
    // ProductClient를 사용하여 외부 상품 정보 서비스로부터 상품 정보를 가져옵니다.
    return s.productClient.FetchProductDetails(productID)
}