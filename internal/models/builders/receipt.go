package builders

import (
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
)

func ReceiptFromGQL(qBuilder *gql.QueryBuilder) (builder *models.Builder){
	if qBuilder == nil {
		return
	}

	builder = &models.Builder{}

	getReceiptFilters(builder, qBuilder.Filter)
	getReceiptPreloads(builder, qBuilder.Load)
	getReceiptJoins(builder, qBuilder.Join)
	builder = getPagination(builder, qBuilder.Pagination)
	builder = getOrder(builder, qBuilder.Order)
	builder = getGroup(builder, qBuilder.Group)

	return
}

func getReceiptFilters(builder *models.Builder, filter map[string]interface{}) {
	if filter == nil {
		return
	}

	for k,v := range filter{
		getReceiptFilter(builder, k, v)
	}

}	

func getReceiptFilter(builder *models.Builder, code string, param interface{} ) {
	where := models.Where{}

	switch(code){
		case "id": 
			where.Condition = "receipts.id = ?"
			where.Params = []interface{} { param }

			break
		case "code":
			where.Condition = "receipts.code = ?"
			where.Params = []interface{} { param }

			break
		default:
			return
	}
	
	builder.Where = append(builder.Where, where)
}

func getReceiptPreloads(builder *models.Builder, param []string) *models.Builder{
	if len(param) == 0{
		return builder
	}

	for _,elem := range param{
		getReceiptPreload(builder,elem)
	}

	return builder
}

func getReceiptPreload(builder *models.Builder, param string) *models.Builder{
	if param == ""{
		return builder
	}

	var preload models.Preload 

	switch(param){
		case "receipt_article":
			preload.Load = "ReceiptArticle"
		default:
			return builder
	}

	builder.Preloads = append(builder.Preloads, preload)
	
	return builder
}

func getReceiptJoins(builder *models.Builder, param []string) *models.Builder{
	if len(param) == 0{
		return builder
	}

	for _,elem := range param{
		getReceiptJoin(builder, elem)		
	}
	
	return builder
}

func getReceiptJoin(builder *models.Builder, param string) *models.Builder{
	if param == ""{
		return builder
	}

	var join string

	switch(param){
		default:
			return builder
	}

	builder.Joins = append(builder.Joins, join)
	
	return builder
}


