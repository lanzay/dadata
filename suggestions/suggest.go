package suggestions

import "github.com/lanzay/dadata/models"
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

