package services

import "delivery_fee_api/models"

func CalculateDeliveryFee(req models.DeliveryFeeRequest) float64 {
    baseFee := 5.0
    weightFee := req.Weight * 0.5
    if weightFee > 10 {
        weightFee = 10
    }
    distanceFee := req.Distance * 0.2
    tempFee := 0.0
    if req.Temperature < 0 || req.Temperature > 30 {
        tempFee = 2.0
    }
    totalFee := baseFee + weightFee + distanceFee + tempFee
    if req.Priority {
        totalFee *= 1.1
    }
    return totalFee
}
