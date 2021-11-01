package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func GetWTables() []*models.WTable{
	return []*models.WTable{
		&models.WTable{
			Name: cons.TableUsers.Name(),
			Code: string(cons.TableUsers),
			Status: true,
			States: GetWStatesUser(),
		},
		&models.WTable{
			Name: cons.TableArticles.Name(),
			Code: string(cons.TableArticles),
			Status: true,
		},
		&models.WTable{
			Name: cons.TablePrices.Name(),
			Code: string(cons.TablePrices),
			Status: true,
		},
		&models.WTable{
			Name: cons.TableReceipts.Name(),
			Code: string(cons.TableReceipts),
			Status: true,
		},
		&models.WTable{
			Name: cons.TableReceiptArticles.Name(),
			Code: string(cons.TableReceiptArticles),
			Status: true,
		},
	}
}


