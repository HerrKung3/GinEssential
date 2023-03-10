package vo

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required"` // name字段是required
}
