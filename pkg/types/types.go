package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.п)
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.п.).
type Category string


//Status ...
type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)


type Payment struct {
	ID int 
	Amount Money
	Category Category
	Status Status
}
