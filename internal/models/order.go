package models

type Order struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

type OrderStruct struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`

	Delivery struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`

	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       uint32 `json:"amount"`
		PaymentDT    uint64 `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost uint32 `json:"delivery_cost"`
		GoodsTotal   uint32 `json:"goods"`
		CustomFee    uint32 `json:"custom_fee"`
	} `json:"payment"`

	Items []struct {
		ChrtId      uint64 `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       uint32 `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        uint8  `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  uint32 `json:"total_price"`
		NmId        uint64 `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`

	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	Shardkey          string `json:"shardkey"`
	SmId              uint   `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}
