package models

type Order struct {
	Id   int         `json:"id" validate:"required"`
	Data OrderStruct `json:"data" validate:"required"`
}

type OrderStruct struct {
	OrderUid    string `json:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Entry       string `json:"entry" validate:"required"`

	Delivery struct {
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone" validate:"required"`
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		Address string `json:"address" validate:"required"`
		Region  string `json:"region" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
	} `json:"delivery" validate:"required"`

	Payment struct {
		Transaction  string `json:"transaction" validate:"required"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency" validate:"required"`
		Provider     string `json:"provider" validate:"required"`
		Amount       uint32 `json:"amount" validate:"required"`
		PaymentDT    uint64 `json:"payment_dt" validate:"required"`
		Bank         string `json:"bank" validate:"required"`
		DeliveryCost uint32 `json:"delivery_cost" validate:"required"`
		GoodsTotal   uint32 `json:"goods"`
		CustomFee    uint32 `json:"custom_fee"`
	} `json:"payment" validate:"required"`

	Items []struct {
		ChrtId      uint64 `json:"chrt_id" validate:"required"`
		TrackNumber string `json:"track_number" validate:"required"`
		Price       uint32 `json:"price" validate:"required"`
		Rid         string `json:"rid"`
		Name        string `json:"name" validate:"required"`
		Sale        uint8  `json:"sale"`
		Size        string `json:"size" validate:"required"`
		TotalPrice  uint32 `json:"total_price" validate:"required"`
		NmId        uint64 `json:"nm_id" validate:"required"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items" validate:"required"`

	Locale            string `json:"locale" validate:"required"`
	InternalSignature string `json:"internal_signature" `
	CustomerId        string `json:"customer_id" validate:"required"`
	DeliveryService   string `json:"delivery_service" validate:"required"`
	Shardkey          string `json:"shardkey"`
	SmId              uint   `json:"sm_id" validate:"required"`
	DateCreated       string `json:"date_created" validate:"required"`
	OofShard          string `json:"oof_shard"`
}
