package Validator

import (
	"fmt"
	"gateway_api/App/Handlers/GORM"
	artranslations "gateway_api/App/Handlers/Validator/validator-translations/ar"
	"gateway_api/Helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

var (
	uni       *ut.UniversalTranslator
	Validator *validator.Validate
	Trans     ut.Translator
	Language  string
)

func NewValidator(c *gin.Context) {
	Validator = validator.New()
	_ = Validator.RegisterValidation("unique", ValidateUnique)
	_ = Validator.RegisterValidation("exists", ValidateExists)
	_ = Validator.RegisterValidation("time", ValidateTime)
	_ = Validator.RegisterValidation("date", ValidateDate)
	_ = Validator.RegisterValidation("regexp", ValidateRegexp)
	_ = Validator.RegisterValidation("strongpass", StrongPassword)
	Language = c.GetHeader("Accept-Language")
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
			tag := f.Tag()
			msg := CustomMessage(tag)
			if msg == "" {
				msg = f.Translate(Trans)
			}
			errs[f.Field()] = []string{msg}
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

func ValidateTime(fl validator.FieldLevel) bool {
	pattern := `(^$|^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$)`
	re := regexp.MustCompile(pattern)
	return re.MatchString(fl.Field().String())
}

func ValidateDate(fl validator.FieldLevel) bool {
	pattern := `^$|\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(fl.Field().String())
}

func ValidateRegexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

func StrongPassword(fl validator.FieldLevel) bool {
	Password := fl.Field().String()
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
