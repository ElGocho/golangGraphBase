package mutations 

import (
	"context"
	"github.com/google/uuid"

	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	sr "sa_web_service/internal/services"

)

func (r *MutationResolver) CreateReceipts(ctx context.Context, input gql.CUReceiptInput) ([]*models.Receipt, error) {
	db := helpers.NewSession(r.DB)

	db = db.Begin()

	receipts, err := sr.CreateReceipts(db, input.Receipts)	

	if err != nil {
		db.Rollback()
		return nil, err
	}

	err = sr.ReceiptsSendEmailToSellers(db,receipts, consts.QueueEmailsEventReceiptCreate, r.MQChannel, r.Queues[consts.QueueEmails])

	if err != nil {
		db.Rollback()
		return nil, err
	}

	err = sr.ReceiptsSendEmailToClients(receipts, consts.QueueEmailsEventReceiptCreate, r.MQChannel, r.Queues[consts.QueueEmails])

	if err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	return receipts, err
}


func (r *MutationResolver) StatusOffer(ctx context.Context, id uuid.UUID, status bool) (*models.Receipt, error) {
	db := helpers.NewSession(r.DB)	
	receipt, err := sr.StatusOffer(db, id, status)	

	return receipt, err
}


