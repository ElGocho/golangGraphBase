package mutations 

import (
	"context"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"

)

func (r *MutationResolver) CreateReceipts(ctx context.Context, input gql.CUReceiptInput) ([]*models.Receipt, error) {
	receipts, err := sr.CreateReceipts(r.DB, input.Receipts)	

	return receipts, err
}


