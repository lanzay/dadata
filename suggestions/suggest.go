package suggestions

import (
	"github.com/lanzay/dadata/models"
	"github.com/kataras/iris/core/errors"
	"log"
	"fmt"
)
/*
Call https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address
*/
func (daData *DaData) SuggestAddresses(query string, count int) ([]models.SuggestionsAddress, error) {

	q := Query{
		Query:query,
		Count:count,
		}
	
	SuggestionsAddress := models.SuggestionsRespond{}
	if sendErr := daData.sendRequest("address", q, &SuggestionsAddress); nil == sendErr {
		return SuggestionsAddress.Suggestions, nil
	} else {
		return nil, sendErr
	}
}

//Main full adress
func (daData *DaData) SuggestAddress(query string, count int) (models.SuggestionsAddress, error) {
	
	x := models.SuggestionsAddress{}
	
	res0, err := daData.SuggestAddresses(query, 3)
	if err != nil {
		log.Println(err)
		return x, errors.New("Adress Error")
	}
	
	fmt.Println("[Unrestricted_value]", res0[0].Unrestricted_value)
	res1, err := daData.SuggestAddresses(res0[0].Unrestricted_value, 1)
	if err != nil {
		log.Println(err)
		return x, errors.New("Adress Error")
	}
	return res1[0], nil
}
