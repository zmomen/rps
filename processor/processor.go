package processor

import (
	"math"
	"rps/model"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Processor struct{}

var (
	// In-memory database of receipts and their points.
	ReceiptPointsDB = map[string]int{}
)

// Define a global regular expression pattern

func (p *Processor) CalculatePoints(request model.ReceiptRequest) model.ReceiptResponse {
	receiptPoints := 0
	receiptPoints += p.calcPointsRetailerName(request.Retailer)
	receiptPoints += p.calcPointsReceiptTotal(request.Total)
	receiptPoints += p.calcPointsItems(request.Items)
	receiptPoints += p.calcPointsDateTime(request.PurchaseDate, request.PurchaseTime)

	// TODO: generate some unique ID per receipt
	resp := model.ReceiptResponse{
		ID: "idd",
	}
	ReceiptPointsDB[resp.ID] = receiptPoints
	return resp
}

func (p *Processor) GetPoints(receiptID string) model.PointsResponse {
	return model.PointsResponse{
		Points: ReceiptPointsDB[receiptID],
	}
}

func (p *Processor) calcPointsRetailerName(retailer string) int {
	totalPoints := 0
	for _, ch := range strings.TrimSpace(retailer) {
		// check if character rune is alphanumeric
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			totalPoints += 1
		}
	}
	return totalPoints
}

func (p *Processor) calcPointsReceiptTotal(total string) int {
	totalPoints := 0
	totalStringArr := strings.Split(total, ".")

	// extract decimal points
	if totalStringArr[1] == "00" {
		totalPoints += 75
	} else if conversion, _ := strconv.Atoi(totalStringArr[1]); conversion%25 == 0 {
		totalPoints += 25
	}

	return totalPoints
}

func (p *Processor) calcPointsItems(items []model.Item) int {
	totalPoints := 0

	// 5 pts for every two items
	totalPoints += 5 * (len(items) / 2)

	// if item length is a multiple of 3, multiply price by 0.2 and round up.
	for _, item := range items {
		itemPointTotal := 0
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			priceFloatVal, _ := strconv.ParseFloat(item.Price, 64)
			itemPointTotal += int(math.Ceil(priceFloatVal * 0.2))
		}
		totalPoints += itemPointTotal
	}

	return totalPoints
}

func (p *Processor) calcPointsDateTime(purchaseDate, purchaseTime string) int {
	totalPoints := 0
	date, _ := time.Parse("2006-01-02 15:04", purchaseDate+" "+purchaseTime)

	// check if odd day
	if date.Day()%2 == 1 {
		totalPoints += 6
	}

	// check if hour is after 2pm and before 4pm
	if date.Hour() >= 14 && date.Minute() > 0 && date.Hour() < 16 {
		totalPoints += 10
	}

	return totalPoints
}
