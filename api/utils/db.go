package utils

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
	"strconv"
)

type ListRequest struct {
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
	Skip     int       `json:"skip"`
	OrderBy  string    `json:"orderBy"`
	Order    string    `json:"order"`
	Filters  *[]Filter `json:"filters"`
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"Operator"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
}

func NewListRequest(r *http.Request, v reflect.Value) *ListRequest {
	var request ListRequest

	request.PageSize = 20
	request.Order = "desc"
	request.OrderBy = "created_At"
	request.Page = 1
	request.FromRequest(r, v)
	return &request
}

func (listRequest *ListRequest) FromRequest(r *http.Request, v reflect.Value) *ListRequest {
	q := r.URL.Query()

	pageSize := q.Get("pageSize")
	page := q.Get("page")
	OrderBy := q.Get("orderBy")
	Order := q.Get("order")

	pageNumber, _ := strconv.Atoi(page)
	if pageNumber > 0 {
		listRequest.Page = pageNumber
	}

	pageSizeLimit, _ := strconv.Atoi(pageSize)
	if pageSizeLimit > 0 {
		listRequest.PageSize = pageSizeLimit
	}

	if OrderBy == "" {
		OrderBy = "created_at"
	}

	if Order == "" || !(Order == "desc" || Order == "asc") {
		Order = "desc"
	}

	listRequest.Order = Order
	listRequest.OrderBy = OrderBy
	listRequest.Skip = (listRequest.Page - 1) * listRequest.PageSize
	listRequest.Filters = ParseFilters(r, v)

	return listRequest
}

func ParseFilters(r *http.Request, v reflect.Value) *[]Filter {

	query := r.URL.Query()
	filters := []Filter{}

	for j := 0; j < v.NumField(); j++ {

		reflectValue := v.Field(j)
		reflectType := reflectValue.Type()
		fieldName := v.Type().Field(j).Tag.Get("json")

		//TODO: should check reflect type if Ptr or Struct to handle nested types and fields but not needed in this task

		fieldQueryValue := query[fieldName]
		if len(fieldQueryValue) > 0 && fieldQueryValue[0] != "" {
			filterObject := CreateFilter(reflectType.String(), fieldName, fieldQueryValue)
			filters = append(filters, filterObject)
		}

	}
	return &filters
}

func CreateFilter(fieldTypeString string, fieldName string, fieldQueryValue []string) Filter {

	if len(fieldQueryValue) == 0 {
		fieldQueryValue = append(fieldQueryValue, "")
	}

	var filterStruct Filter
	filterStruct.Field = fieldName
	filterStruct.Operator = FILTER_OPERATOR_EQUAL
	filterStruct.Type = fieldTypeString

	switch fieldTypeString {

	case "string":
		filterStruct.Value = fieldQueryValue[0]
		break
	case "int":
		i, _ := strconv.Atoi(fieldQueryValue[0])
		filterStruct.Value = i
		break
	case "int64":
		i, _ := strconv.ParseInt(fieldQueryValue[0], 10, 64)
		filterStruct.Value = i
		break
	case "float64":
		s, _ := strconv.ParseFloat(fieldQueryValue[0], 64)
		filterStruct.Value = s
		break
	case "bool":
		b, _ := strconv.ParseBool(fieldQueryValue[0])
		filterStruct.Value = b
		break

	default:
		//application.GetLogger().Error("can' set filter to field of type --> ", fieldTypeString)
	}

	return filterStruct

}

func GenerateSqlCondition(db *gorm.DB, filters *[]Filter) *gorm.DB {

	for _, filter := range *filters {
		if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "string" {
			db = db.Where("\""+filter.Field+"\"::text"+" ILIKE ?", "%"+filter.Value.(string)+"%")

		} else if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "float64" {
			db = db.Where("\""+filter.Field+"\""+" = ?", filter.Value)

		} else if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "int" {
			db = db.Where("\""+filter.Field+"\""+" = ?", filter.Value)

		} else if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "int64" {
			db = db.Where("\""+filter.Field+"\""+" = ?", filter.Value)

		} else if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "bool" {
			db = db.Where("\""+filter.Field+"\""+" = ?", filter.Value)

		} else if filter.Operator == FILTER_OPERATOR_EQUAL && filter.Type == "uuid" {
			db = db.Where("\""+filter.Field+"\""+" = ?", filter.Value)

		} else {
			db = db.Where("\""+filter.Field+"\"::text"+" ILIKE ?", "%"+filter.Value.(string)+"%")
		}
	}
	return db
}
