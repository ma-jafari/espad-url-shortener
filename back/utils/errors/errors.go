package utils

var ErrorMessages = map[int]string{
	ERROR_INVALID_PARAMETERS: "پارامتر های ارسالی نامعتبر",
	ERROR_NOT_FOUND:          "موردی یافت نشد",
	INTERNAL_SERVER_ERROR:    "خطای سرور رخ داده است",
}

const (
	ERROR_INVALID_PARAMETERS = 400
	ERROR_NOT_FOUND          = 404
	INTERNAL_SERVER_ERROR    = 500
)
