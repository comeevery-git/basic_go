package dto

// 상품정보 응답값 객체
type ProductDetails struct {
    ID          int     `json:"id"`
    ProductName string  `json:"product_name"`
    ProductDescription string  `json:"product_description,omitempty"`
    Price       float64 `json:"price"`
    UseYn bool          `json:"use_yn"`
}
