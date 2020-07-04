package captcha

import (
	"github.com/burhon94/goCaptchaMojo/base64MyCaptchaDigit"
	"github.com/burhon94/goCaptchaMojo/modeles"
)

func GenerateCaptcha(postParameters modeles.ConfigJsonBody) (base64blob, captchaId string) {
	var config modeles.ConfigDigit
	config = postParameters.DigitParams
	captchaId, captchaInterfaceInstance := base64MyCaptchaDigit.GenerateCaptcha(postParameters.Id, config)
	base64blob = base64MyCaptchaDigit.CaptchaWriteToBase64Encoding(captchaInterfaceInstance)

	return
}

func VerifyCaptcha(postParameters modeles.ConfigJsonBody) (ResultOK bool){
	verifyResult := base64MyCaptchaDigit.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)

	return verifyResult
}

