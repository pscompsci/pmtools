package forms

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type MultipartForm struct {
	Errors errors
}

func NewMultipartForm() *MultipartForm {
	return &MultipartForm{
		errors(map[string][]string{}),
	}
}

func (f *MultipartForm) Required(r *http.Request, fields ...string) {
	r.ParseMultipartForm(32 << 20)
	for _, field := range fields {
		_, _, err := r.FormFile(field)
		if err != nil {
			f.Errors.Add(field, "the field cannot be empty")
		}
	}
}

func (f *MultipartForm) IsType(r *http.Request, field, extension string) {
	r.ParseMultipartForm(32 << 20)
	_, headers, err := r.FormFile(field)
	if err != nil {
		return
	}
	ext := filepath.Ext(headers.Filename)
	if ext != extension {
		f.Errors.Add(field, fmt.Sprintf("invalid file type. expected '%s', given '%s'", extension, ext))
	}
}

func (f *MultipartForm) Valid() bool {
	return len(f.Errors) == 0
}
