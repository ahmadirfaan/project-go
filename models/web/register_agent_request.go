package web

type RegisterAgentRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
	AgentName   string `json:"agentName" validate:"required"`
	NoHandphone string `json:"noHandphone" validate:"required,min=9,max=12,numeric"`
	DistrictId  string `json:"districtId" validate:"required,len=7,numeric"`
	Address     string `json:"address" validate:"required"`
}
