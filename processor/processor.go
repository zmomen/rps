package processor

import "rps/model"

type Processor struct{}


var (
	// In-memory database of receipts and their points.
	ReceiptPointsDB = map[string]float64{}
)
func (p *Processor) CalculatePoints() model.ReceiptResponse {
	
	// TODO: do some calculation and generate some ID for receipt
	
	return model.ReceiptResponse{
		ID: "something",
	}
}

func (p *Processor) GetPoints(receiptID string) model.PointsResponse {
	return model.PointsResponse{
		Points: ReceiptPointsDB[receiptID],
	}
}
