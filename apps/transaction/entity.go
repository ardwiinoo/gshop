package transaction

import (
	"encoding/json"
	"time"

	"github.com/ardwiinoo/online-shop/infra/response"
)

type TransactionStatus uint8

const (
	TransactionStatus_Created = 1
	TransactionStatus_Progress = 10
	TransactionStatus_InDelivery = 15
	TransactionStatus_Completed = 20

	TRX_CREATED = "CREATED"
	TRX_PROGRESS = "PROGRESS"
	TRX_IN_DELIVERY = "IN DELIVERY"
	TRX_COMPLETED = "COMPLETED"
	TRX_UNKNOWN = "UNKNOWN STATUS"
)

var (
	MappingTransactionStatus = map[TransactionStatus]string{
		TransactionStatus_Created: TRX_CREATED,
		TransactionStatus_Progress: TRX_PROGRESS,
		TransactionStatus_InDelivery: TRX_IN_DELIVERY,
		TransactionStatus_Completed: TRX_COMPLETED,
	}
)

type Transaction struct {
	Id           int     		 	`db:"id"`
	UserPublicId        string  		 	`db:"user_public_id"`
	ProductId    uint    		 	`db:"product_id"`
	ProductPrice uint    		 	`db:"product_price"`
	Amount       uint8   		 	`db:"amount"`
	SubTotal     uint    		 	`db:"sub_total"`
	PlatformFee  uint    		 	`db:"platform_fee"`
	GrandTotal   uint    		 	`db:"grand_total"`
	Status	   	 TransactionStatus  `db:"status"`
	ProductJSON  json.RawMessage 	`db:"product_snapshot"`
	CreatedAt 	 time.Time          `db:"created_at"`
	UpdatedAt 	 time.Time          `db:"updated_at"`
} 

func NewTransaction(userPublicId string) Transaction {
	return Transaction{
		UserPublicId: userPublicId,
		Status: TransactionStatus_Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewTransactionFromCreateRequest(req CreateTransactionRequestPayload) Transaction {
	return Transaction{
		UserPublicId: req.UserPublicId,
		Amount: req.Amount,
		Status: TransactionStatus_Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t Transaction) Validate() (err error) {
	if t.Amount == 0 {
		return response.ErrAmountInvalid
	}

	return
}

func (t Transaction) ValidateStock(productStock uint8) (err error) {
	if t.Amount > productStock {
		return response.ErrAmountGreaterThanStock
	}
	
	return
}

func (t *Transaction) SetSubTotal() {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}
}

func (t *Transaction) SetPlatformFee(platformFee uint) (*Transaction) {
	t.PlatformFee = platformFee

	return t
}

// set subtotal and grandtotal
func (t *Transaction) SetGrandTotal() (*Transaction) {
	if t.GrandTotal == 0 {
		t.SetSubTotal()
		
		t.GrandTotal = t.SubTotal + t.PlatformFee
	}

	return t
}

// set productId, price, and json
func (t *Transaction) FromProduct(product Product) (*Transaction) {
	t.ProductId = uint(product.Id)
	t.ProductPrice = uint(product.Price)

	t.SetProductJSON(product)

	return t
}

func (t *Transaction) SetProductJSON(product Product) (err error) {
	productJSON, err := json.Marshal(product)

	if err != nil {
		return
	}

	t.ProductJSON = productJSON

	return
}

func (t Transaction) GetProduct() (product Product, err error) {
	err = json.Unmarshal(t.ProductJSON, &product)

	if err != nil {
		return
	}

	return
}

func (t Transaction) GetStatus() string {
	status, ok := MappingTransactionStatus[t.Status]

	if !ok {
		return TRX_UNKNOWN
	}

	return status
}