package model

type User struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	memo *string `json:"memo,omitempty"`
	UseYn bool `json:"use_yn"`
}

type Seller struct {
	ID int `json:"id"`
	SellerName string `json:"seller_name"`
	SellerEmail string `json:"seller_email"`
	memo *string `json:"memo,omitempty"`
	UseYn bool `json:"use_yn"`
}


/**
type Product struct {
	ID int `json:"id"`
	ProductName string `json:product_name`
	ProductDescription *string `json:product_description,omitempty` // omitempty: empty value 일 때 변환 대상에서 제외시킴. 단, nil pointer 일 때임
	Price float64 `json:price`
	Count int `json:count`
	UseYn bool
	// java 에서의 보다 정확한 소수점 계산을 위해 BigDecimal 을 사용하듯
	// Go에서는 정확한 소수점 계산을 위해 float64 대신 shopspring/decimal 라이브러리를 사용할 수는 있다
}
*/