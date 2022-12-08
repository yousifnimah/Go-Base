package Validator

func CustomMessage(tag string) string {
	messages := map[string]map[string]string{
		"ar": {
			"strongpass": "يجب أن تحتوي على أحرف صغيرة وأحرف كبيرة وحالة خاصة ورقم",
			"unique":     "يجب ان تكون قيمة الحقل غير مستخدمة من قبل",
			"exists":     "القيمة المُدخلة غير موجودة",
			"time":       "يجب ان تكون القيمة بصيغة الوقت HH:ii:ss",
			"date":       "يجب ان تكون القيمة بصيغة تاريخ yyyy-mm-dd",
			"regexp":     "الصيغة غير صحيحة",
		},
		"en": {
			"strongpass": "Must contain lowercase, uppercase, spacial-case, and number",
			"unique":     "The field value must not be used before",
			"exists":     "The value entered does not exist",
			"time":       "The value must be in the time format HH:ii:ss",
			"date":       "The value must be in the date format YYYY-mm-dd",
			"regexp":     "The format is incorrect",
		},
	}
	return messages[Language][tag]
}
