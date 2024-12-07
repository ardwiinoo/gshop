package transaction

import "encoding/json"

type Transaction struct {
	Id           int     		 `db:"id"`
	Email        string  		 `db:"email"`
	ProductId    uint    		 `db:"product_id"`
	ProductPrice uint    		 `db:"product_price"`
	Amount       uint8   		 `db:"amount"`
	SubTotal     uint    		 `db:"sub_total"`
	PlatformFee  uint    		 `db:"platform_fee"`
	GrandTotal   uint    		 `db:"grand_total"`
	Status	   	 uint8  		 `db:"status"`
	ProductJSON  json.RawMessage `db:"product_snapshot"`
	CreatedAt 	 string          `db:"created_at"`
	UpdatedAt 	 string          `db:"updated_at"`
} 