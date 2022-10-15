package web

type CreateTransactionRequest struct {
	TransactionTypeId uint   `json:"transactionTypeId" validate:"required"`
	AgentId           uint   `json:"agentId" validate:"required"`
	Address           string `json:"address" validate:"required"`
	DistrictId        string `json:"districtId" validate:"required,len=7,numeric"`
	Amount            uint64 `json:"amount" validate:"required"`
}
