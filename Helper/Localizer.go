package Helper

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var EnLocalizer *i18n.Localizer
var ARLocalizer *i18n.Localizer

func Localize(c *gin.Context, MessageID string) string {
	var Localizer *i18n.Localizer
	Language := c.GetHeader("Accept-Language")
	localizeConfig := i18n.LocalizeConfig{
		MessageID: MessageID,
	}
	switch Language {
	case "ar":
		Localizer = ARLocalizer
	default:
		Localizer = EnLocalizer
	}
	localization, _ := Localizer.Localize(&localizeConfig)
	return localization
}

func InitLocalize() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("Resources/langs/en.toml")
	bundle.MustLoadMessageFile("Resources/langs/ar.toml")
	EnLocalizer = i18n.NewLocalizer(bundle, "en", "en")
	ARLocalizer = i18n.NewLocalizer(bundle, "ar", "ar")
}
