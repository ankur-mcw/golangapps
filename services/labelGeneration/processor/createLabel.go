package processor

import (
	"errors"
	"log"
)

type LabelGenerationRequest struct {
    CarrierMoniker  string `json:"carrier_moniker"`
    RetailerMoniker string `json:"retailer_moniker"`
}

type LabelGenerationResponse struct {
    Message string `json:"message"`
}

func GenerateLabel(request LabelGenerationRequest) (*LabelGenerationResponse, error) {

    response := &LabelGenerationResponse{}
    carrierMoniker := request.CarrierMoniker
    if len(carrierMoniker) == 0 {
        return response, errors.New("Carrier moniker cannot be empty")
    }
    log.Printf("CarrierMoniker=%s", carrierMoniker)

    retailerMoniker := request.RetailerMoniker
    if len(retailerMoniker) == 0 {
        return response, errors.New("Retailer moniker cannot be empty")
    }
    log.Printf("RetailerMoniker=%s", retailerMoniker)
    return &LabelGenerationResponse{"created"}, nil
}
