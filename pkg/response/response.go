package response

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidPayload     = "invalid payload request"
	InvalidQuery       = "invalid query request"
	InternalServer     = "internal server error"
	SomethingWentWrong = "something went wrong"
	Unauthorized       = "unauthorized request"
)

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"status"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
