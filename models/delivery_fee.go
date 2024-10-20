package models

type DeliveryFeeRequest struct {
    Weight      float64 `json:"weight"`
    Distance    float64 `json:"distance"`
    Priority    bool    `json:"priority"`
    Temperature float64 `json:"temperature"`
}

type DeliveryFeeResponse struct {
    DeliveryFee float64 `json:"delivery_fee"`
}
