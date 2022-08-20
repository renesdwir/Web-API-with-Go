package dataInput

type DataInput struct {
	Id         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	GradeClass string `json:"grade_class" binding:"required"`
}