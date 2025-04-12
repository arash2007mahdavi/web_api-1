package helper

import "github.com/arash2007mahdavi/web-api-1/api/validations"

type BaseResponse struct {
	Result          any                            `json:"result,omitempty"`
	Success         bool                           `json:"success,omitempty"`
	ValidationError *[]validations.ValidationError `json:"validationError,omitempty"`
	Error           any                            `json:"error,omitempty"`
}

func GenerateBaseResponse(result any, success bool) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
	}
}

func GenerateBaseResponseWithError(success bool, err error) *BaseResponse {
	return &BaseResponse{
		Success:    success,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(success bool, err error) *BaseResponse {
	return &BaseResponse{
		Success:         success,
		ValidationError: validations.GetValidationErrors(err),
	}
}
