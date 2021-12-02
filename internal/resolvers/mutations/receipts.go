package mutations 

import (
	"context"
	"github.com/google/uuid"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"

)

func (r *MutationResolver) CreateReceipts(ctx context.Context, input gql.CUReceiptInput) ([]*models.Receipt, error) {
	receipts, err := sr.CreateReceipts(r.DB, input.Receipts)	

	return receipts, err
}


func (r *MutationResolver) StatusOffer(ctx context.Context, id uuid.UUID, status bool) (*models.Receipt, error) {
	receipt, err := sr.StatusOffer(r.DB, id, status)	

	return receipt, err
}


