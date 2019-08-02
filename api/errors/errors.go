package errors

import "fmt"

type APIError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
} 

var (
	InternalServerError = APIError{500, "SERVERERR", "Internal Server Error", "An unkown error occured"}
	PlanNotFound        = APIError{400, "PLANNOTFND", "Plan Not Found", "No plan was found matching the queried id"}
	NoPlansFound        = APIError{400, "NOPLANSFND", "No Plans Found", "There were no plans found"}
	InvalidForm         = APIError{400, "INVALIDFORM", "Invalid Form", "The submitted form is not valid"}
	EmailOrPassword     = APIError{400, "EMAILPASSMISS", "Email or password missing", "The Email or password missing"}
	EmailNotFound       = APIError{400, "EMAILNOTFND", "Email Not Found", "No matching email address was found"}
	WrongPassword       = APIError{401, "WRONGPASS", "Wrong Password", "password is incorrect"}
	EmailTaken          = APIError{400, "EMAILTAKEN", "Email Taken", "email already exists"}
	AuthCookieMissing   = APIError{401, "AUTHCOOKMISS", "Auth Cookie Missing", "Auth Cookie is missing"}
	RefresCookieMissing = APIError{401, "REFCOOKMISS", "Refresh Cookie Missing", "Refresh Cookie is missing"}
	TokenInvalid        = APIError{401, "TOKENINVALID", "Token Invalid", "Authorisation token invalid"}
	InvalidIdentifier   = APIError{401, "INVIDENTIFIER", "Invlaid identifier", "Invlaid identifier"}
	CSRFTokenInvalid    = APIError{401, "CSRFTOKEN", "CSRF Token", "CSRF token is invalid"}
)

func (err *APIError) Error() string {
	return fmt.Sprintf("Code: %d, Title: %s, Detail: %s", err.Code, err.Title, err.Detail)
}
