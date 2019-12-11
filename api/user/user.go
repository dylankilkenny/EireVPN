package user

import (
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/integrations/stripe"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"fmt"
	"net/http"
	"strconv"
	"time"

	stripego "github.com/stripe/stripe-go"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser verifies a users details are correct, returning a jwt token to the user
func LoginUser(c *gin.Context) {
	var userLogin models.User
	var userDb models.User

	if err := c.BindJSON(&userLogin); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/login - LoginUser()",
			Code:  errors.EmailOrPassword.Code,
			Extra: map[string]interface{}{"Email": userLogin.Email},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.EmailOrPassword.Status, errors.EmailOrPassword)
		return
	}

	userDb.Email = userLogin.Email
	if err := userDb.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/login - LoginUser()",
			Code:  errors.EmailNotFound.Code,
			Extra: map[string]interface{}{"Email": userLogin.Email},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.EmailNotFound.Status, errors.EmailNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/login - LoginUser()",
			Code:  errors.WrongPassword.Code,
			Extra: map[string]interface{}{"Email": userLogin.Email},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.WrongPassword.Status, errors.WrongPassword)
		return
	}

	var usersession models.UserAppSession
	if err := usersession.New(userDb.ID); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/login - LoginUser() - Create session",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": userDb.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	authToken, refreshToken, csrfToken, err := jwt.Tokens(usersession)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:   "/login - LoginUser()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": userDb.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	conf := config.GetConfig()

	// TODO: Change the domain name and add correct maxAge time
	authCookieMaxAge := 15 * 60 // 15 minutes in seconds
	c.SetCookie("authToken", authToken, authCookieMaxAge, "/", conf.App.Domain, false, true)

	// TODO: Change the domain name and add correct maxAge time
	refreshCookieMaxAge := 24 * 60 * 60 // 72 hours in seconds
	c.SetCookie("refreshToken", refreshToken, refreshCookieMaxAge, "/", conf.App.Domain, false, true)

	c.Header("X-CSRF-Token", csrfToken)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   gin.H{"firstname": userDb.FirstName},
	})
}

// SignUpUser registers a new user
func SignUpUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/signup - SignUpUser()",
			Code:  errors.InvalidForm.Code,
			Extra: map[string]interface{}{"Email": user.Email},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	var userTmp models.User
	userTmp.Email = user.Email
	if err := userTmp.Find(); err == nil {
		logger.Log(logger.Fields{
			Loc:   "/signup - SignUpUser()",
			Code:  errors.EmailTaken.Code,
			Extra: map[string]interface{}{"Email": user.Email},
			Err:   errors.EmailTaken.Detail,
		})
		c.JSON(errors.EmailTaken.Status, errors.EmailTaken)
		return
	}

	user.Type = models.UserTypeNormal
	if err := user.Create(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/signup - SignUpUser()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"Email": user.Email},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   make([]string, 0),
	})
}

func StripeSession(c *gin.Context) {
	var user models.User

	userID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/user/session/:planid - StripeSession()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
	}
	user.ID = userID.(uint)
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/session/:planid - StripeSession()",
			Code: errors.UserNotFound.Code,
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": errors.UserNotFound.Detail,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	planID, _ := strconv.ParseUint(c.Param("planid"), 10, 64)
	var plan models.Plan
	plan.ID = uint(planID)

	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/session/:planid - StripeSession()",
			Code: errors.PlanNotFound.Code,
			Extra: map[string]interface{}{
				"PlanID": c.Param("planid"),
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	if user.StripeCustomerID == "" {
		customer, err := user.CreateStripeCustomer()
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "/user/session/:planid - StripeSession()",
				Code: errors.StripeCreateCustomerErr.Code,
				Extra: map[string]interface{}{
					"UserID": userID,
					"Detail": errors.StripeCreateCustomerErr.Detail,
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.StripeCreateCustomerErr.Status, errors.StripeCreateCustomerErr)
			return
		}
		user.StripeCustomerID = customer.ID
		if err := user.Save(); err != nil {
			logger.Log(logger.Fields{
				Loc:  "/user/session/:planid - CreateSession()",
				Code: errors.StripeCreateCustomerErr.Code,
				Extra: map[string]interface{}{
					"StripeCustomerID": customer.ID,
					"UserID":           userID,
					"Detail":           "Error saving stripe customer ID to user",
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
			return
		}

	}
	var sessionID string
	if plan.PlanType == models.PlanTypeSubscription {
		stripeSession, err := stripe.CreateSubscriptionSession(plan.StripePlanID, user.StripeCustomerID, fmt.Sprint(user.ID))
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "/user/session/:planid - CreateSession()",
				Code: errors.StripeCreateSessionErr.Code,
				Extra: map[string]interface{}{
					"StripePlanID": plan.StripePlanID,
					"Detail":       errors.StripeCreateSessionErr.Detail,
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.StripeCreateSessionErr.Status, errors.StripeCreateSessionErr)
			return
		}
		sessionID = stripeSession.ID
	}
	if plan.PlanType == models.PlanTypePayAsYouGo {
		var cart models.Cart
		cart.UserID = user.ID
		cart.PlanID = plan.ID
		if err := cart.Save(); err != nil {
			logger.Log(logger.Fields{
				Loc:  "/user/session/:planid - CreateSession()",
				Code: errors.StripeCreateCustomerErr.Code,
				Extra: map[string]interface{}{
					"cartID": cart.ID,
					"UserID": userID,
					"Detail": "Error saving cart",
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
			return
		}

		stripeSession, err := stripe.CreatePAYGSession(plan.Name, user.StripeCustomerID, cart.ID, plan.Amount)
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "/user/session/:planid - CreateSession()",
				Code: errors.StripeCreateSessionErr.Code,
				Extra: map[string]interface{}{
					"StripePlanID": plan.StripePlanID,
					"Detail":       errors.StripeCreateSessionErr.Detail,
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.StripeCreateSessionErr.Status, errors.StripeCreateSessionErr)
			return
		}
		sessionID = stripeSession.ID
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   map[string]interface{}{"session_id": sessionID},
	})
}

func StripeUpdatePaymentSession(c *gin.Context) {
	var user models.User

	userID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/user/updatepayment- StripeUpdatePaymentSession()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
	}
	user.ID = userID.(uint)
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/updatepayment- StripeUpdatePaymentSession()",
			Code: errors.UserNotFound.Code,
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": errors.UserNotFound.Detail,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	customer, err := stripe.GetCustomer(user.StripeCustomerID)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/updatepayment- StripeUpdatePaymentSession()",
			Code: errors.StripeCustomerNotFound.Code,
			Extra: map[string]interface{}{
				"StripeCustomerID": user.StripeCustomerID,
				"UserID":           user.ID,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.StripeCustomerNotFound.Status, errors.StripeCustomerNotFound)
		return
	}
	var subcscription stripego.Subscription
	for _, item := range customer.Subscriptions.Data {
		if item.Status == stripego.SubscriptionStatusActive {
			subcscription = *item
		}
	}
	stripeSetupSession, err := stripe.CreateSessionSetup(customer.ID, subcscription.ID)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/updatepayment- StripeUpdatePaymentSession()",
			Code: errors.StripeCreateSessionSetupErr.Code,
			Extra: map[string]interface{}{
				"StripeCustomerID": customer.ID,
				"StripeSubID":      subcscription.ID,
				"Detail":           errors.StripeCreateSessionSetupErr.Detail,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.StripeCreateSessionSetupErr.Status, errors.StripeCreateSessionSetupErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   map[string]interface{}{"session_id": stripeSetupSession.ID},
	})
}

func Webhook(c *gin.Context) {
	conf := config.GetConfig()
	webhookevent, err := stripe.WebhookEventHandler(c.Request.Body, c.Request.Header.Get("Stripe-Signature"), conf.Stripe.EndpointSecret)
	if err != nil {
		logger.Log(logger.Fields{
			Loc: "/user/webhook - Webhook()",
			Extra: map[string]interface{}{
				"Detail": "Error fetching webhook event",
			},
			Err: err.Error(),
		})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	switch webhookevent.Type {
	case "checkout.session.completed":
		if webhookevent.CheckoutModeSubscription {
			var plan models.Plan
			plan.StripePlanID = webhookevent.StripePlanID
			if err := plan.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.PlanNotFound.Code,
					Extra: map[string]interface{}{
						"PlanID": webhookevent.StripePlanID,
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
				return
			}
			var userPlan models.UserPlan
			userPlan.UserID = webhookevent.UserID
			userPlan.PlanID = plan.ID
			userPlan.Active = true
			userPlan.StartDate = time.Now()
			userPlan.ExpiryDate = time.Unix(webhookevent.StripeSubscriptionEndPeriod, 0)
			if err := userPlan.Save(); err != nil {
				logger.Log(logger.Fields{
					Loc:   "/user/webhook - Webhook()",
					Code:  errors.InternalServerError.Code,
					Extra: map[string]interface{}{"UserID": userPlan.UserID},
					Err:   err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}
		}

		if webhookevent.CheckoutModePayment {
			var cart models.Cart
			cart.ID = webhookevent.CartID
			if err := cart.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.InternalServerError.Code,
					Extra: map[string]interface{}{
						"CartID": webhookevent.CartID,
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}

			var plan models.Plan
			plan.ID = cart.PlanID
			if err := plan.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.PlanNotFound.Code,
					Extra: map[string]interface{}{
						"PlanID": cart.PlanID,
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
				return
			}

			var userPlan models.UserPlan
			userPlan.UserID = cart.UserID
			userPlan.PlanID = plan.ID
			userPlan.Active = true
			userPlan.StartDate = time.Now()
			userPlan.ExpiryDate = time.Now().Add(time.Hour * time.Duration(plan.IntervalCount))
			if err := userPlan.Save(); err != nil {
				logger.Log(logger.Fields{
					Loc:   "/user/webhook - Webhook()",
					Code:  errors.InternalServerError.Code,
					Extra: map[string]interface{}{"UserPlanID": userPlan.UserID},
					Err:   err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}

			if err := cart.Delete(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.InternalServerError.Code,
					Extra: map[string]interface{}{
						"CartID": cart.UserID,
						"Detail": "Failed to delete cart",
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}
		}

	case "invoice.payment_succeeded":
		// We only want to continue if the invoice type is
		// from a recurring subscription payment rather
		// than an invoice from the subscription creation.
		if webhookevent.InvoiceTypeSubscription {
			var plan models.Plan
			plan.StripePlanID = webhookevent.StripePlanID
			if err := plan.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.PlanNotFound.Code,
					Extra: map[string]interface{}{
						"PlanID": webhookevent.StripePlanID,
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
				return
			}
			var user models.User
			user.StripeCustomerID = webhookevent.StripeCustomerID
			if err := user.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.UserNotFound.Code,
					Extra: map[string]interface{}{
						"CustomerID": webhookevent.StripeCustomerID,
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
				return
			}
			var userPlan models.UserPlan
			userPlan.UserID = user.ID
			userPlan.PlanID = plan.ID
			if err := userPlan.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:  "/user/webhook - Webhook()",
					Code: errors.InternalServerError.Code,
					Extra: map[string]interface{}{
						"UserID": userPlan.UserID,
						"Detail": "Could not find user_plan record",
					},
					Err: err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}

			userPlan.Active = true
			userPlan.ExpiryDate = time.Unix(webhookevent.StripeSubscriptionEndPeriod, 0)
			if err := userPlan.Save(); err != nil {
				logger.Log(logger.Fields{
					Loc:   "/user/webhook - Webhook()",
					Code:  errors.InternalServerError.Code,
					Extra: map[string]interface{}{"UserID": userPlan.UserID},
					Err:   err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{})
}

func CancelSubscription(c *gin.Context) {
	var user models.User
	userID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/user/updatepayment- StripeUpdatePaymentSession()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
	}
	user.ID = userID.(uint)
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/updatepayment- StripeUpdatePaymentSession()",
			Code: errors.UserNotFound.Code,
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": errors.UserNotFound.Detail,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	subscription, err := stripe.CustomerSubscription(user.StripeCustomerID)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/cancel - CancelSubscription()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"UserID": user.ID,
				"Detail": "Error fetching customer subscription",
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	var plan models.Plan
	plan.StripePlanID = subscription.Plan.ID
	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/webhook - Webhook()",
			Code: errors.PlanNotFound.Code,
			Extra: map[string]interface{}{
				"PlanID": subscription.Plan.ID,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	var userPlan models.UserPlan
	userPlan.UserID = user.ID
	userPlan.PlanID = plan.ID
	if err := userPlan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/webhook - Webhook()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"UserID": userPlan.UserID,
				"Detail": "Could not find user_plan record",
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	if err := userPlan.Delete(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/webhook - Webhook()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"UserPlanID": userPlan.ID,
				"Detail":     "Failed to delete user plan",
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	err = stripe.CancelSubscription(subscription.ID)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:  "/user/cancel - CancelSubscription()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"UserID": user.ID,
				"Detail": "Error canceling customer subscription",
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// // ChangePasswordRequest sends the user a link to change their password
// func ChangePasswordRequest(c *gin.Context) {
// 	db := db.GetDB()
// 	email := c.PostForm("email")
// 	var user models.User

// 	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
// 		c.AbortWithStatusJSON(errors.EmailNotFound.Status, errors.EmailNotFound)
// 		return
// 	}

// 	// Send email here

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 200,
// 		"errors": make([]string, 0),
// 		"data":   make([]string, 0),
// 	})
// }

// // ChangePassword will authenticate the users token and change their password
// func ChangePassword(c *gin.Context) {
// 	db := db.GetDB()
// 	email := c.PostForm("email")
// 	var user models.User

// 	// _, err := jwt.Validate(c)
// 	// if err != nil {
// 	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 	// 		"status": 401,
// 	// 		"errors": gin.H{
// 	// 			"title":  "Invalid Token",
// 	// 			"detail": "Token provided in auth header is not valid",
// 	// 		},
// 	// 		"data": make([]string, 0),
// 	// 	})
// 	// 	return
// 	// }

// 	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"status": 404,
// 			"errors": gin.H{
// 				"title":  "Email Not Found",
// 				"detail": "No matching email address was found",
// 			},
// 			"data": make([]string, 0),
// 		})
// 		return
// 	}

// 	// Send email here

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 200,
// 		"errors": make([]string, 0),
// 		"data":   make([]string, 0),
// 	})
// }
