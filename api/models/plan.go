package models

import (
	"fmt"
	"time"

	"eirevpn/api/integrations/stripe"

	"github.com/jinzhu/gorm"
)

type AllPlans []Plan

type PlanType string

var (
	PlanTypePayAsYouGo   PlanType = "PAYG"
	PlanTypeSubscription PlanType = "SUB"
	PlanTypeFreeTrial    PlanType = "FREE"
)

// Plan holds the detilas for a given vpn plan on offer
type Plan struct {
	BaseModel
	Name            string   `json:"name" binding:"required"`
	Amount          int64    `json:"amount" binding:"required"`
	Interval        string   `json:"interval" binding:"required"`
	IntervalCount   int64    `json:"interval_count" binding:"required"`
	PlanType        PlanType `json:"plan_type" binding:"required"`
	Currency        string   `json:"currency" binding:"required"`
	StripePlanID    string   `json:"stripe_plan_id"`
	StripeProductID string   `json:"stripe_product_id"`
}

// BeforeCreate sets the CreatedAt column to the current time
func (p *Plan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (p *Plan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (p *Plan) Find() error {
	if err := db().Where(&p).First(&p).Error; err != nil {
		return err
	}
	return nil
}

func (ap *AllPlans) FindAll() error {
	if err := db().Find(&ap).Error; err != nil {
		return err
	}
	return nil
}

func (p *Plan) Create() error {
	if p.PlanType == PlanTypeSubscription {
		if err := p.createStripePlan(); err != nil {
			return err
		}
	}
	if err := db().Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Plan) Delete() error {
	if p.PlanType == PlanTypeSubscription {
		if err := p.deleteStripePlan(); err != nil {
			return err
		}
	}
	if err := db().Delete(&p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Plan) Save() error {
	if p.PlanType == PlanTypeSubscription {
		if err := p.updateStripePlan(); err != nil {
			return err
		}
	}
	if err := db().Save(&p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Plan) createStripePlan() error {
	stripePlanID, stripeProductID, err := stripe.CreatePlan(p.Amount, p.IntervalCount, p.Interval, p.Name, p.Currency)
	if stripePlanID != nil && stripeProductID != nil {
		p.StripePlanID = *stripePlanID
		p.StripeProductID = *stripeProductID
	}
	return err
}

func (p *Plan) updateStripePlan() error {
	return stripe.UpdatePlan(p.StripeProductID, p.Name)
}

func (p *Plan) deleteStripePlan() error {
	return stripe.DeletePlan(p.StripePlanID, p.StripeProductID)
}

// BeforeCreate sets the CreatedAt column to the current time
func (p *Plan) String() string {
	return fmt.Sprintf(
		"ID: %d, Name: %s, Amount: %s, Interval: %d, IntervalCount: %d, Currency: %d",
		p.ID,
		p.Name,
		p.Amount,
		p.Interval,
		p.IntervalCount,
		p.Currency,
	)
}
