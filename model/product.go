package model

// em aspas o nome de retorno dos dados
type Product struct {
	ID        int     `json:"id_product"`
	Name      string  `json:"product_name"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}
