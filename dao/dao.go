package dao

type GetAnimalRes struct {
	ID      string `json:"_id"`
	Species string `json:"species,omitempty"`
}

type CreateAnimalReq struct {
	Species string `json:"species,omitempty"`
}

type CreateAnimalRes struct {
	ID      string `json:"_id"`
	Species string `json:"species,omitempty"`
}

type DeleteAnimalReq struct {
	Species string `json:"species,omitempty"`
}
