package types

type CommentData struct {
	Data []Comment `json:"data"`
}

type Comment struct {
	Comment        string `json:"comment"`
	Name           string `json:"name"`
	DescScore      int    `json:"desc_score"`
	LogisticsScore int    `json:"logistics_score"`
	ServiceScore   int    `json:"service_score"`
}
