package handlers

import (
    "encoding/json"
    "net/http"
    "delivery_fee_api/models"
    "delivery_fee_api/services"
    "delivery_fee_api/utils"
)

func HandleDeliveryFeeRequest(w http.ResponseWriter, r *http.Request) {
    var req models.DeliveryFeeRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    feeChannel := make(chan float64)
    go func() {
        fee := services.CalculateDeliveryFee(req)
        feeChannel <- fee
    }()

    deliveryFee := <-feeChannel
    response := models.DeliveryFeeResponse{
        DeliveryFee: deliveryFee,
    }
    utils.RespondWithJSON(w, http.StatusOK, response)
}
