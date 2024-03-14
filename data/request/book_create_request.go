package request

type BookCreateRequest struct {
	Name        *string `validate:"required min=1,max=100" json:"name"`
	Description *string `validate:"required min=1,max=200" json:"description"`
	Author      *string `validate:"required min=1,max=100" json:"author"`
	Image       *string `validate:"required min=1,max=100" json:"image"`
	Genre       *string `validate:"required min=1,max=100" json:"genre"`
	PublicDate  *string `validate:"required min=1,max=100" json:"public_date"`
}
