package stripe

import (
	"eirevpn/api/config"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentmethod"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/setupintent"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/webhook"
)

var conf config.Config

type WebhookEvent struct {
	Type                        string
	CheckoutModeSubscription    bool
	CheckoutModePayment         bool
	InvoiceTypeSubscription     bool
	StripeSubscriptionEndPeriod int64
	StripePlanID                string
	StripeCustomerID            string
	UserID                      uint
	CartID                      uint
}

func Init() {
	conf = config.GetConfig()
	stripe.Key = conf.Stripe.SecretKey
}

func CreatePlan(amount, intervalCount int64, interval, name, currency string) (*string, *string, error) {
	if conf.Stripe.IntegrationActive {
		params := &stripe.PlanParams{
			Amount:        &amount,
			Interval:      &interval,
			IntervalCount: &intervalCount,
			Product: &stripe.PlanProductParams{
				Name: &name,
			},
			Currency: &currency,
		}
		stripePlan, err := plan.New(params)
		return &stripePlan.ID, &stripePlan.Product.ID, err
	}
	return nil, nil, nil
}

func UpdatePlan(StripeProductID, name string) error {
	if conf.Stripe.IntegrationActive {
		_, err := product.Update(StripeProductID, &stripe.ProductParams{
			Name: &name,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func DeletePlan(StripePlanID, StripeProductID string) error {
	if conf.Stripe.IntegrationActive {
		_, err := plan.Del(StripePlanID, nil)
		if err != nil {
			return err
		}
		_, err = product.Del(StripeProductID, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSessionSubscription(planID, customerID, userID string) (*stripe.CheckoutSession, error) {
	if conf.Stripe.IntegrationActive {
		params := &stripe.CheckoutSessionParams{
			Customer:          stripe.String(customerID),
			ClientReferenceID: stripe.String(userID),
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
				Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
					&stripe.CheckoutSessionSubscriptionDataItemsParams{
						Plan: stripe.String(planID),
					},
				},
			},
			SuccessURL: stripe.String(conf.Stripe.SuccessURL),
			CancelURL:  stripe.String(conf.Stripe.ErrorURL),
		}

		return session.New(params)
	}
	return nil, nil
}

func CreateSessionPAYG(planName, customerID string, cartID uint, planAmount int64) (*stripe.CheckoutSession, error) {
	if conf.Stripe.IntegrationActive {
		params := &stripe.CheckoutSessionParams{
			Customer:          stripe.String(customerID),
			ClientReferenceID: stripe.String(string(cartID)),
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				&stripe.CheckoutSessionLineItemParams{
					Name:     stripe.String(planName),
					Amount:   stripe.Int64(planAmount),
					Currency: stripe.String(string(stripe.CurrencyEUR)),
					Quantity: stripe.Int64(1),
				},
			},
			SuccessURL: stripe.String(conf.Stripe.SuccessURL),
			CancelURL:  stripe.String(conf.Stripe.ErrorURL),
		}

		return session.New(params)
	}
	return nil, nil
}

func CreateSessionSetup(customerID, subscriptionID string) (*stripe.CheckoutSession, error) {
	if conf.Stripe.IntegrationActive {
		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			SetupIntentData: &stripe.CheckoutSessionSetupIntentDataParams{
				Params: stripe.Params{
					Metadata: map[string]string{
						"customer_id":     customerID,
						"subscription_id": subscriptionID,
					},
				},
			},
			Mode:       stripe.String(string(stripe.CheckoutSessionModeSetup)),
			SuccessURL: stripe.String(conf.Stripe.SuccessURL + "?setup=1"),
			CancelURL:  stripe.String(conf.Stripe.ErrorURL + "?setup=1"),
		}
		return session.New(params)
	}
	return nil, nil
}

func CreateCustomer(customerEmail, firstName, lastName string, userID uint) (*stripe.Customer, error) {
	if conf.Stripe.IntegrationActive {
		params := &stripe.CustomerParams{
			Name:        stripe.String(firstName + " " + lastName),
			Email:       stripe.String(customerEmail),
			Description: stripe.String("Customer for " + customerEmail),
		}
		params.AddMetadata("EireVPN_UserID", string(userID))
		return customer.New(params)
	}
	return nil, nil
}

func GetSubscription(subscriptionId string) (*stripe.Subscription, error) {
	if conf.Stripe.IntegrationActive {
		return sub.Get(subscriptionId, nil)
	}
	return nil, nil
}

func GetCustomer(customerId string) (*stripe.Customer, error) {
	if conf.Stripe.IntegrationActive {
		return customer.Get(customerId, nil)
	}
	return nil, nil
}

func GetSetupIntent(setupIntentID string) (*stripe.SetupIntent, error) {
	if conf.Stripe.IntegrationActive {
		return setupintent.Get(setupIntentID, nil)
	}
	return nil, nil
}

func AddPaymentMethodToCustomer(customerID, paymentMethodID string) error {
	paymentMethodParams := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	}
	_, err := paymentmethod.Attach(paymentMethodID, paymentMethodParams)
	return err
}

func SetCustomerDefaultPaymentMethod(customerID, paymentMethodID string) error {
	customerParams := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(paymentMethodID),
		},
	}
	_, err := customer.Update(customerID, customerParams)
	return err
}

func UpdateCustomerPaymentMethod(customerID, paymentMethodID string) error {
	err := AddPaymentMethodToCustomer(customerID, paymentMethodID)
	if err != nil {
		return err
	}
	err = SetCustomerDefaultPaymentMethod(customerID, paymentMethodID)
	if err != nil {
		return err
	}
	return nil
}

func WebhookEventHandler(body io.ReadCloser, stripeSignature, endpointSecret string) (*WebhookEvent, error) {
	var webhookEvent WebhookEvent
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	event, err := webhook.ConstructEvent(payload, stripeSignature, endpointSecret)
	if err != nil {
		return nil, err
	}

	webhookEvent.Type = event.Type
	switch event.Type {
	case "checkout.session.completed":
		var checkoutSession stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &checkoutSession)
		if err != nil {
			return nil, err
		}

		if checkoutSession.Mode == stripe.CheckoutSessionModePayment {
			webhookEvent.CheckoutModePayment = true
			cartID, _ := strconv.ParseUint(checkoutSession.ClientReferenceID, 10, 64)
			webhookEvent.CartID = uint(cartID)
		}

		if checkoutSession.Mode == stripe.CheckoutSessionModeSetup {
			setupIntent, err := GetSetupIntent(checkoutSession.SetupIntent.ID)
			if err != nil {
				return nil, err
			}

			err = UpdateCustomerPaymentMethod(setupIntent.Metadata["customer_id"], setupIntent.PaymentMethod.ID)
			if err != nil {
				return nil, err
			}
			webhookEvent.CheckoutModeSubscription = false
		}

		if checkoutSession.Mode == stripe.CheckoutSessionModeSubscription {
			var planID string
			for _, item := range checkoutSession.DisplayItems {
				planID = item.Plan.ID
			}
			webhookEvent.StripePlanID = planID
			sub, err := GetSubscription(checkoutSession.Subscription.ID)
			if err != nil {
				return nil, err
			}
			webhookEvent.StripeSubscriptionEndPeriod = sub.CurrentPeriodEnd
			webhookEvent.CheckoutModeSubscription = true
			userID, _ := strconv.ParseUint(checkoutSession.ClientReferenceID, 10, 64)
			webhookEvent.UserID = uint(userID)
		}

	case "invoice.payment_succeeded":
		var invoice stripe.Invoice
		err := json.Unmarshal(event.Data.Raw, &invoice)
		if err != nil {
			return nil, err
		}
		if invoice.BillingReason == stripe.InvoiceBillingReasonSubscription {
			sub, err := GetSubscription(invoice.Subscription)
			if err != nil {
				return nil, err
			}
			webhookEvent.StripePlanID = sub.Plan.ID
			webhookEvent.StripeCustomerID = invoice.Customer.ID
			webhookEvent.InvoiceTypeSubscription = true
		}
	}
	return &webhookEvent, nil
}
