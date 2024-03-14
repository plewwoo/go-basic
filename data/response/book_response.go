package response

type BookReponse struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Author      *string `json:"author"`
	Image       *string `json:"image"`
	Genre       *string `json:"genre"`
	PublicDate  *string `json:"public_date"`
}
