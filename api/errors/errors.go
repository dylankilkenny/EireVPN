package errors

import "fmt"

type APIError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type InternalError struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var (
	ConfigLoadError             = InternalError{"CONFIGLOAD", "Failed to load config", "An error occured when trying to load the configuration file"}
	InternalServerError         = APIError{500, "SERVERERR", "Internal Server Error", "An unkown error occured"}
	StripeCreatePlanErr         = APIError{500, "STRIPECREATEPLAN", "Stripe Create Plan Error", "Failed to create plan with stripe"}
	StripeCreateSessionErr      = APIError{500, "STRIPECREATESESS", "Stripe Create Session Error", "Failed to create session with stripe"}
	StripeCreateSessionSetupErr = APIError{500, "STRIPECREATESESSSET", "Stripe Create Session Setup Error", "Failed to create session setup with stripe"}
	StripeDeletePlanErr         = APIError{500, "STRIPEDELPLAN", "Stripe Delete Plan Error", "Failed to delete plan with stripe"}
	StripeDeleteProductErr      = APIError{500, "STRIPEDELPROD", "Stripe Delete Product Error", "Failed to delete product with stripe"}
	StripeUpdatePlanErr         = APIError{500, "STRIPEUPDPLAN", "Stripe Update Plan Error", "Failed to update plan with stripe"}
	StripeUpdateProductErr      = APIError{500, "STRIPEUPDPROD", "Stripe Update Product Error", "Failed to update product with stripe"}
	StripeCreateCustomerErr     = APIError{500, "STRIPECREACUSTERR", "Stripe Create Customer Error", "Failed to create customer with stripe"}
	StripeUpdatePayMethodErr    = APIError{400, "STRIPEUDATEPAY", "Stripe Update Payment Method Error", "Failed to update customer payment method with stripe"}
	StripeCustomerNotFound      = APIError{400, "STRCUSTNOTFOUND", "Stripe Customer Not Found", "Failed to find customer with stripe"}
	PlanNotFound                = APIError{400, "PLANNOTFND", "Plan Not Found", "No plan was found matching the queried id"}
	ServerNotFound              = APIError{400, "CONNNOTFND", "Server Not Found", "No server was found matching the queried id"}
	NoPlansFound                = APIError{400, "NOPLANSFND", "No Plans Found", "There were no plans found"}
	InvalidForm                 = APIError{400, "INVALIDFORM", "Invalid Form", "The submitted form is not valid"}
	EmailOrPassword             = APIError{400, "EMAILPASSMISS", "Email or password missing", "The Email or password missing"}
	EmailNotFound               = APIError{400, "EMAILNOTFND", "Email Not Found", "No matching email address was found"}
	UserNotFound                = APIError{400, "USERNOTFND", "User Not Found", "No matching user found for the supplied ID"}
	WrongPassword               = APIError{401, "WRONGPASS", "Wrong Password", "password is incorrect"}
	EmailTaken                  = APIError{400, "EMAILTAKEN", "Email Taken", "email already exists"}
	AuthCookieMissing           = APIError{401, "AUTHCOOKMISS", "Auth Cookie Missing", "Auth Cookie is missing"}
	RefresCookieMissing         = APIError{401, "REFCOOKMISS", "Refresh Cookie Missing", "Refresh Cookie is missing"}
	TokenInvalid                = APIError{401, "TOKENINVALID", "Token Invalid", "Authorisation token invalid"}
	InvalidIdentifier           = APIError{401, "INVIDENTIFIER", "Invlaid identifier", "Invlaid identifier"}
	CSRFTokenInvalid            = APIError{401, "CSRFTOKEN", "CSRF Token", "CSRF token is invalid"}
	ProtectedRouted             = APIError{400, "PROTECTROUTE", "Protected Route", "You do not have the correct permissions to access this route."}
)

func (err *APIError) Error() string {
	return fmt.Sprintf("Code: %d, Title: %s, Detail: %s", err.Code, err.Title, err.Detail)
}
