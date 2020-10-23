package structs

type Pagination struct {
	Message  string `json:"message"`
	Page     uint   `json:"page" form:"page"`
	PageSize uint   `json:"page_size" form:"page_size"`
	Total    uint   `json:"total" form:"total"`
}
