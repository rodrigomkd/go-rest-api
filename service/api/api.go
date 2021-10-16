package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
)

type IApiService interface {
	GetActivities() []model.Activity
}

type ApiService struct {
	activitiesUri string
	api           IApiService
}

func New(activitiesUri string) *ApiService {
	return &ApiService{
		activitiesUri: activitiesUri,
	}
}

//ReadAPI - call api and returns the response
func readAPI(uri string) []byte {
	response, err := http.Get(uri)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func (api ApiService) GetActivities() []model.Activity {
	responseData := readAPI(api.activitiesUri)
	var responseObject []model.Activity
	json.Unmarshal(responseData, &responseObject)

	return responseObject
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
