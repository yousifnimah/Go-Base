package Validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"base/App/Handlers/GORM"
	artranslations "base/App/Handlers/Validator/validator-translations/ar"
	"base/Helper"
	"reflect"
	"strings"
)

var (
	uni       *ut.UniversalTranslator
	Validator *validator.Validate
	Trans     ut.Translator
)

func NewValidator(c *gin.Context) {
	Validator = validator.New()
	_ = Validator.RegisterValidation("unique", ValidateUnique)
	_ = Validator.RegisterValidation("exists", ValidateExists)
	Language := c.GetHeader("Accept-Language")
	registerLanguage(Validator, Language)

	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Validate(err error) map[string][]string {
	if err != nil {
		errs := make(map[string][]string)
		errorsList := err.(validator.ValidationErrors)
		for _, f := range errorsList {
			errs[f.Field()] = []string{f.Translate(Trans)}
		}
		return errs
	}
	return nil
}

func registerLanguage(validate *validator.Validate, Language string) {
	EnLocale := en.New()
	uni = ut.New(EnLocale, EnLocale)
	Trans, _ = uni.GetTranslator("en")
	if Language == "ar" {
		_ = artranslations.RegisterDefaultTranslations(validate, Trans)
	} else {
		_ = entranslations.RegisterDefaultTranslations(validate, Trans)
	}
}

func ValidateUnique(fl validator.FieldLevel) bool {
	TableName := fl.Param()
	FieldName := fl.FieldName()

	//custom field name
	FieldParts := strings.Split(fl.Param(), ":")
	if len(FieldParts) > 1 {
		TableName = FieldParts[0]
		FieldName = FieldParts[1]
	}

	FieldType := fl.Field().Type().String()
	Value := ""

	if FieldType == "int" {
		Value = fmt.Sprintf("%d", fl.Field().Int())
	} else {
		Value = fl.Field().String()
	}

	Value = Helper.EscapeSChars(Value)

	var Results int
	db := GORM.OpenConnection()
	db.Raw(fmt.Sprintf("select count(id) from %s where %s = '%s'", TableName, FieldName, Value)).Scan(&Results)
	GORM.CloseConnection(db)
	return Results == 0
}

func ValidateExists(fl validator.FieldLevel) bool {
	TableName := fl.Param()
	FieldName := fl.FieldName()

	//custom field name
	FieldParts := strings.Split(fl.Param(), ":")
	if len(FieldParts) > 1 {
		TableName = FieldParts[0]
		FieldName = FieldParts[1]
	}

	FieldType := fl.Field().Type().String()
	Value := ""

	if FieldType == "int" {
		Value = fmt.Sprintf("%d", fl.Field().Int())
	} else {
		Value = fl.Field().String()
	}

	Value = Helper.EscapeSChars(Value)

	var Results int
	db := GORM.OpenConnection()
	db.Raw(fmt.Sprintf("select count(id) from %s where %s = '%s'", TableName, FieldName, Value)).Scan(&Results)
	GORM.CloseConnection(db)
	return Results > 0
}
