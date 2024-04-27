package email

type Template int

const (
	OTPCode Template = iota
)

func SendEmail(template Template, to string)
