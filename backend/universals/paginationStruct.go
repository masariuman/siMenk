package universals

type Pagination struct {
	Current_page int         `json:"current_page"`
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	Per_page     int         `json:"per_page"`
}
