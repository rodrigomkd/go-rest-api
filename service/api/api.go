package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
)

type IHttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type ApiService struct {
	activitiesUri string
	httpClient    IHttpClient
}

func New(activitiesUri string, client IHttpClient) *ApiService {
	return &ApiService{
		activitiesUri: activitiesUri,
		httpClient:    client,
	}
}

func (api ApiService) GetActivities() ([]model.Activity, error) {
	response, err := api.httpClient.Get(api.activitiesUri)
	log.Println("API response: ", response)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	var responseObject []model.Activity
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func convertActivitiesCsv(activities []model.Activity) [][]string {
	var records = [][]string{}

	for i := 0; i < len(activities); i++ {
		activity := activities[i]
		var record = []string{}
		record = append(record, strconv.Itoa(activity.Id))
		record = append(record, activity.Title)
		record = append(record, activity.DueDate)
		record = append(record, strconv.FormatBool(activity.Completed))

		records = append(records, record)
	}

	return records
}
