package types

// Money предоставляет денежную сумму в минимальных единицач {центы, копейки, дирамы и т.д.}
type Money int64

// Category представляет собой категорию в которой был совершен платеж
type Category string

//Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
	Category Category
}
