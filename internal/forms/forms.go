package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Form create a custom form struct, embed url values object
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New init form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required adalah function untuk cek bahwa field pada form tidak bolaeh kosong
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "field ini tidak boleh kosong!")
		}
	}
}

//MinLength adalah func untuk validasi panjang karakter pada form field
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("field ini minimal harus %d karakter", length))
		return false
	}
	return true
}

//IsEmail cek apakah email input form is valid or not
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "invalid email address")
	}
}

// //Has cek form required field kosong apa tidak
// func (f *Form) Has(field string, r *http.Request) bool {
// 	x := r.Form.Get(field)
// 	if x == "" {
// 		f.Errors.Add(field, "field tidak boleh kosong")
// 		return false
// 	}
// 	return true
// }
