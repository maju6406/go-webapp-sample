package dto

import (
	"encoding/json"

	"github.com/maju6406/go-webapp-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

const (
	required string = "required"
	max      string = "max"
	min      string = "min"
)

// RegBookDto defines a data transfer object for register.
type RegBookDto struct {
	Title      string `validate:"required,min=3,max=50" json:"title"`
	Isbn       string `validate:"required,min=10,max=20" json:"isbn"`
	CategoryID uint   `json:"categoryId"`
	FormatID   uint   `json:"formatId"`
}

// NewRegBookDto is constructor.
func NewRegBookDto() *RegBookDto {
	return &RegBookDto{}
}

// Create creates a book model from this DTO.
func (b *RegBookDto) Create() *model.Book {
	return model.NewBook(b.Title, b.Isbn, b.CategoryID, b.FormatID)
}

// Validate performs validation check for the each item.
func (b *RegBookDto) Validate() map[string]string {
	return validateDto(b)
}

func validateDto(b interface{}) map[string]string {
	result := make(map[string]string)
	err := validator.New().Struct(b)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "ID":
					switch errors[i].Tag() {
					case required:
						result["id"] = "書籍IDが存在しません"
					}
				case "Title":
					switch errors[i].Tag() {
					case required, min, max:
						result["title"] = "書籍タイトルは、3文字以上50文字以下で入力してください"
					}
				case "Isbn":
					switch errors[i].Tag() {
					case required, min, max:
						result["isbn"] = "ISBNは、10文字以上20文字以下で入力してください"
					}
				}
			}
		}
		return result
	}
	return nil
}

// ToString is return string of object
func (b *RegBookDto) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}

// ChgBookDto defines a data transfer object for changes or updates.
type ChgBookDto struct {
	ID         uint   `validate:"required" json:"id"`
	Title      string `validate:"required,gte=3,lt=50" json:"title"`
	Isbn       string `validate:"required,gte=10,lt=20" json:"isbn"`
	CategoryID uint   `json:"categoryId"`
	FormatID   uint   `json:"formatId"`
}

// NewChgBookDto is constructor.
func NewChgBookDto() *ChgBookDto {
	return &ChgBookDto{}
}

// Validate performs validation check for the each item.
func (b *ChgBookDto) Validate() map[string]string {
	return validateDto(b)
}

// ToString is return string of object
func (b *ChgBookDto) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
