package request

// Number godoc
type Number struct {
	Count int `json:"count" validate:"number"`
}
