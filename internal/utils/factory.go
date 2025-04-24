package utils

import "fs.seaotterms.com-backend/internal/dto"

func InitResponse() dto.CommonResponse {
	return dto.CommonResponse{
		Message: "",
		Data:    nil,
	}
}
