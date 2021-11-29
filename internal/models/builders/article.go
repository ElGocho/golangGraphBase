package builders

import (
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"fmt"
)

func ArticleFromGQL(qBuilder *gql.QueryBuilder) (builder *models.Builder){
	if qBuilder == nil {
		return
	}

	builder = &models.Builder{}

	getArticleFilters(builder, qBuilder.Filter)
	getArticlePreloads(builder, qBuilder.Load)
	getArticleJoins(builder, qBuilder.Join)
	builder = getPagination(builder, qBuilder.Pagination)
	builder = getOrder(builder, qBuilder.Order)
	builder = getGroup(builder, qBuilder.Group)

	return
}

func getArticleFilters(builder *models.Builder, filter map[string]interface{}) {
	if filter == nil {
		return
	}

	for k,v := range filter{
		getArticleFilter(builder, k, v)
	}

}	

func getArticleFilter(builder *models.Builder, code string, param interface{} ) {
	where := models.Where{}

	switch(code){
		case "id": 
			where.Condition = "articles.id = ?"
			where.Params = []interface{} { param };

			break
		case "name":
			where.Condition = "articles.name like ?"
			where.Params = []interface{} { fmt.Sprintf("%c%v%c",'%',param,'%')}

			break
		default:
			return
	}
	
	builder.Where = append(builder.Where, where)
}

func getArticlePreloads(builder *models.Builder, param []string) *models.Builder{
	if len(param) == 0{
		return builder
	}

	for _,elem := range param{
		getArticlePreload(builder,elem)
	}

	return builder
}

func getArticlePreload(builder *models.Builder, param string) *models.Builder{
	if param == ""{
		return builder
	}

	var preload models.Preload 

	switch(param){
		case "prices":
			preload.Load = "Prices"
		case "seller":
			preload.Load = "Seller"
		default:
			return builder
	}

	builder.Preloads = append(builder.Preloads, preload)
	
	return builder
}

func getArticleJoins(builder *models.Builder, param []string) *models.Builder{
	if len(param) == 0{
		return builder
	}

	for _,elem := range param{
		getArticleJoin(builder, elem)		
	}
	
	return builder
}

func getArticleJoin(builder *models.Builder, param string) *models.Builder{
	if param == ""{
		return builder
	}

	var join string

	switch(param){
		case "seller":
			join = "Seller"
		default:
			return builder
	}

	builder.Joins = append(builder.Joins, join)
	
	return builder
}


