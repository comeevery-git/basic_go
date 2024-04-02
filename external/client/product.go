package client

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/m/external/dto"
)

// ProductClient는 외부 상품 서비스 API를 호출하는 클라이언트를 정의합니다.
type ProductClient struct {
    baseURL string // 외부 서비스의 기본 URL
}

// NewProductClient는 ProductClient 인스턴스를 생성하고 반환합니다.
func NewProductClient(baseURL string) *ProductClient {
    return &ProductClient{baseURL: baseURL}
}

// FetchProductDetails는 주어진 상품 ID에 대한 상품 정보를 외부 서비스로부터 가져옵니다.
func (pc *ProductClient) FetchProductDetails(productID int) (*dto.ProductDetails, error) {
    // 외부 API 호출
    resp, err := http.Get(pc.baseURL + "/products/" + strconv.Itoa(productID))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var productDetails dto.ProductDetails
    if err := json.NewDecoder(resp.Body).Decode(&productDetails); err != nil {
        return nil, err
    }

    return &productDetails, nil
}