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
	ReadWorkers() []model.Worker
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

// with Worker pools
func (cs CSVWorkerService) concuRSwWP(f *os.File, typ string, items int, itemsPerWork int) []model.Worker {
	fcsv := csv.NewReader(f)
	rs := make([]model.Worker, 0)
	numWps := items / itemsPerWork
	jobs := make(chan []string, numWps)
	res := make(chan model.Worker)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- model.Worker) {
		for {
			select {
			case job, ok := <-jobs: // must check for readable state of the channel.
				if !ok {
					return
				}
				results <- parseWorker(job)
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output
			defer wg.Done()
			worker(jobs, res)
		}()
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

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		//items_per_worker
		rs = append(rs, r)
	}

	return rs
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
