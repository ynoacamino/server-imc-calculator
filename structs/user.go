package structs

type User struct {
	ID       int     `json:"user_id"`
	Peso     float32 `json:"peso"`
	Talla    float32 `json:"talla"`
	Response float32 `json:"response"`
}
