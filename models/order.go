package models

import (
	"concurrency-experimentals/utils"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

const (
	OrderStatus_Started           = "STARTED"
	OrderStatus_WaitingForPayment = "WAITING_FOR_PAYMENT"
	OrderStatus_Paid              = "PAID"
	OrderStatus_Canceled          = "CANCELED"
	OrderStatus_Rejected          = "REJECTED"
	OrderStatus_LeftForDelivery   = "LEFT_FOR_DELIVERY"
	OrderStatus_Delivered         = "DELIVERED"
	OrderStatus_Completed         = "COMPLETED"
)

const (
	PaymentType_NotPaid    = "NOT_PAID"
	PaymentType_CreditCard = "CREDIT_CARD"
	PaymentType_DebitCard  = "DEBIT_CARD"
	PaymentType_InCash     = "IN_CASH"
)

const (
	RejectedType_InvalidPaymentType   = "INVALID_PAYMENT_TYPE"
	RejectedType_InvalidOrderQuantity = "PAYMENT_EXPIRED"
	RejectedType_ProductUnavailable   = "PRODUCT_UNAVAILABLE"
)

func GetAllOrderStatus() []string {
	return []string{
		OrderStatus_Started,
		OrderStatus_WaitingForPayment,
		OrderStatus_Paid,
		OrderStatus_Canceled,
		OrderStatus_Rejected,
		OrderStatus_LeftForDelivery,
		OrderStatus_Delivered,
		OrderStatus_Completed,
	}
}

func GetAllPaymentTypes() []string {
	return []string{
		PaymentType_CreditCard,
		PaymentType_DebitCard,
		PaymentType_InCash,
	}
}

func GetAllRejectedType() []string {
	return []string{
		RejectedType_InvalidPaymentType,
		RejectedType_InvalidOrderQuantity,
		RejectedType_ProductUnavailable,
	}
}

type Order struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	ProductID    string    `json:"product_id"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	UnitPrice    float64   `json:"unit_price"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
	ReasonReject string    `json:"reason_reject"`
	Paid         int8      `json:"paid"`
	PaymentType  string    `json:"payment_type"`
	IsDeleted    int8      `json:"is_deleted"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	CanceledAt   *time.Time `json:"canceled_at"`
}

func (o *Order) GenerateID() string {
	o.ID = bson.NewObjectId().Hex()
	return o.ID
}

func (o *Order) SetPaid(paid bool) {
	o.Paid = utils.BoolToInt(paid)
}

func (o *Order) SetIsDeleted(isDeleted bool) {
	o.IsDeleted = utils.BoolToInt(isDeleted)
}

func (o *Order) GeneratePaymentType(r *rand.Rand) string {
	o.PaymentType = PaymentType_NotPaid
	if utils.IntToBool(o.Paid) {
		paymentTypes := GetAllPaymentTypes()
		o.PaymentType = paymentTypes[r.Intn(len(paymentTypes))]
	}
	return o.PaymentType
}

func (o *Order) GenerateUpdatedAt(r *rand.Rand) time.Time {
	updatedAt := o.CreatedAt.AddDate(0, r.Intn(12), r.Intn(28))
	o.UpdatedAt = &updatedAt
	return *o.UpdatedAt
}

func (o *Order) GenerateDescription() string {
	b, err := json.Marshal(o)
	if err != nil {
		o.Description = err.Error()
	} else {
		o.Description = string(b)
	}
	return o.Description
}

func (o *Order) ToMongoOrderModel() *MongoOrderModel {
	return &MongoOrderModel{
		ID:           o.ID,
		UserID:       o.UserID,
		ProductID:    o.ProductID,
		Description:  o.Description,
		Quantity:     o.Quantity,
		UnitPrice:    o.UnitPrice,
		Amount:       o.Amount,
		Status:       o.Status,
		ReasonReject: o.ReasonReject,
		Paid:         utils.IntToBool(o.Paid),
		PaymentType:  o.PaymentType,
		IsDeleted:    utils.IntToBool(o.IsDeleted),
		CreatedAt:    *o.CreatedAt,
		UpdatedAt:    *o.UpdatedAt,
		DeletedAt:    *o.DeletedAt,
		CanceledAt:   *o.CanceledAt,
	}
}

type MongoOrderModel struct {
	ID           string    `bson:"_id"`
	UserID       string    `bson:"user_id"`
	ProductID    string    `bson:"product_id"`
	Description  string    `bson:"description"`
	Quantity     int       `bson:"quantity"`
	UnitPrice    float64   `bson:"unit_price"`
	Amount       float64   `bson:"amount"`
	Status       string    `bson:"status"`
	ReasonReject string    `bson:"reason_reject"`
	Paid         bool      `bson:"paid"`
	PaymentType  string    `bson:"payment_type"`
	IsDeleted    bool      `bson:"is_deleted"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	DeletedAt    time.Time `bson:"deleted_at"`
	CanceledAt   time.Time `json:"canceled_at"`
}
