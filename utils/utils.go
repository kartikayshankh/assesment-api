package utils

type ErrorHandler struct {
	DevMessage string                 `json:"developerMessage,omitempty"`
	Request    map[string]interface{} `json:"request,omitempty"`
	Response   map[string]interface{} `json:"response,omitempty"`
	UserId     string                 `json:"userId,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Method     string                 `json:"method,omitempty"`
	Code       int                    `json:"code,omitempty"`
	Status     string                 `json:"status,omitempty"`
}

const (
	SUCCESS                       = "success"
	FAILED                        = "failed"
	DATA_MISSING                  = "data Missing"
	DATA_NOT_FOUND                = "data Not Found"
	EMPLOYEE_CREATED_SUCCESSFULLY = "employee created successfully"
	EMPLOYEE_UPDATED_SUCCESSFULLY = "employee updated successfully"
	EMPLOYEE_DELETED_SUCCESSFULLY = "employee deleted successfully"
	SOMETHING_WENT_WRONG          = "something went wrong"
)

func ErrorMapper(message string) string {
	mapper := map[string]string{
		"ent: employee not found": DATA_NOT_FOUND,
	}
	return mapper[message]
}
