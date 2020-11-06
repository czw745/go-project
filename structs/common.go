package structs

type Pagination struct {
	Page     int         `json:"page" form:"page"`
	PageSize int         `json:"page_size" form:"page_size"`
	Total    int64       `json:"total" form:"total"`
	Data     interface{} `json:"data"`
}

type Response struct {
	Message string `json:"message" form:"message"`
}

type RoleSelect struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
