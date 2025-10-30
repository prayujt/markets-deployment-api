package models

type MarketDeployRequest struct {
	QuestionID    string `json:"questionId"`
	ConditionID   string `json:"conditionId"`
	PositionIDYes string `json:"positionIdYes"`
	PositionIDNo  string `json:"positionIdNo"`
}
