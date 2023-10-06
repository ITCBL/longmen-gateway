package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"longmen-gateway/global"
	myValidator "longmen-gateway/validator"
	"reflect"
	"strings"
)

func InitTrans(locale string) {
	// 修改gin框架中的验证器引擎属性以实现自定义
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义方法以获取JSON标记，并将错误提示中的struct字段替换为相应的JSON字段
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()              // 中文翻译器
		enT := en.New()              // 英文翻译器
		uni := ut.New(enT, zhT, enT) // 第一个参数是备份语言环境，下面的参数是应该支持的语言环境
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			panic(fmt.Errorf("uni.GetTranslator(%s)", locale))
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}

	return
}

// RegisterValidation 自定义验证器
func RegisterValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_mobile", myValidator.ValidateMobile)
		_ = v.RegisterTranslation("valid_mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("valid_mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("valid_mobile", fe.Field())
			return t
		})
	}
}
