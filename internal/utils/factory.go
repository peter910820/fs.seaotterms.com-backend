package utils

import "fs.seaotterms.com-backend/internal/dto"

func InitResponse() dto.Response {
	return dto.Response{
		Message: "",
		Data:    nil,
	}
}
