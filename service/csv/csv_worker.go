package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rodrigomkd/go-rest-api/model"
)

type ICSVWService interface {
	ReadWorkers(typ string, items int, itemsPerWork int) []model.Worker
}

type CSVWorkerService struct {
	csv            ICSVWService
	dataSourcePath string
}

func NewWorker(dataSourcePath string) *CSVWorkerService {
	return &CSVWorkerService{
		dataSourcePath: dataSourcePath,
	}
}

func (cs CSVWorkerService) ReadWorkers(typ string, items int, itemsPerWork int) []model.Worker {
	f, _ := os.Open(cs.dataSourcePath)
	defer f.Close()

	ts := time.Now()
	activities := cs.concuRSwWP(f, typ, items, itemsPerWork)
	te := time.Now().Sub(ts)

	fmt.Println("END Concurrency: ", te)

	return activities
}

//concuRSwWP - Worker pools
func (cs CSVWorkerService) concuRSwWP(f *os.File, typ string, items int, itemsPerWork int) []model.Worker {
	fcsv := csv.NewReader(f)
	rs := make([]model.Worker, items)
	jobs := make(chan []string, items)
	res := make(chan model.Worker)

	wg := new(sync.WaitGroup)

	// start up workers
	for w := 1; w <= items; w++ {
		wg.Add(1)
		go itemsPerWorker(jobs, res, wg, itemsPerWork)
	}

	go func() {
		count := 1
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}

			if rStr[4] == typ {
				jobs <- rStr

				log.Println("rStr: ", rStr)
				if count == items {
					break
				}
				count++
			}
		}

		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	// Now collect all the results...
	// But first, make sure close the result channel when everything was processed
	go func() {
		wg.Wait()
		close(res)
	}()

	index := 0
	for r := range res {
		rs[index] = r
		index++
	}

	return rs
}

func itemsPerWorker(jobs <-chan []string, results chan<- model.Worker, wg *sync.WaitGroup, itemsPerWork int) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	i := 0
	for j := range jobs {
		if i <= itemsPerWork {
			results <- parseWorker(j)
		}
	}
}

func parseWorker(data []string) model.Worker {
	id, _ := strconv.Atoi(data[0])
	completed, _ := strconv.ParseBool(data[3])

	return model.Worker{
		Id:        id,
		Title:     data[1],
		DueDate:   data[2],
		Completed: completed,
		Type:      data[4],
	}
}
