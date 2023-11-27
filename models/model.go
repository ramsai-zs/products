package models

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brandName"`
	Details   string `json:"details"`
	ImageUrl  string `json:"imageUrl"`
}

type Response struct {
	Product Product   `json:"product,omitempty"`
	Variant []Variant `json:"variant,omitempty"`
}

type FinalResponse struct {
	Product    Product       `json:"product,omitempty"`
	SubVariant []VariantResp `json:"variant,omitempty"`
	Variant
}

type ProductResponse struct {
	Product []Product     `json:"product"`
	Variant []VariantResp `json:"variant,omitempty"`
}

type VariantResp struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Details string `json:"details,omitempty"`
}

type Variant struct {
	ID        string `json:"id"`
	ProductID string `json:"ProductID"`
	Name      string `json:"name"`
	Details   string `json:"details"`
}

type Products struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	BrandName string    `json:"brandName"`
	Details   string    `json:"details"`
	ImageURL  string    `json:"imageURL"`
	Variants  []Variant `json:"variants"`
}

type Filters struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
	VariantID   int    `json:"variantId"`
}
