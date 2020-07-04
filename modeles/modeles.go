package modeles

type ConfigDigit struct {
	// Height png height in pixel.
	Height int `json:"height"`
	// Width Captcha png width in pixel.
	Width int `json:"width"`
	// DefaultLen Default number of digits in captcha solution.
	CaptchaLen int `json:"captcha_len"`
	// MaxSkew max absolute skew factor of a single digit.
	MaxSkew float64 `json:"max_skew"`
	// DotCount Number of background circles.
	DotCount int `json:"dot_count"`
}

//configJsonBody json request body.
type ConfigJsonBody struct {
	Id          string      `json:"id"`
	VerifyValue string      `json:"verifyValue"`
	DigitParams ConfigDigit `json:"digit_params"`
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload"`
}
