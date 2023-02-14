package weather

import (
	// Standard Libs
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	// Third Party Libs
	cron "github.com/go-co-op/gocron"
)

func Stop(job *cron.Job, s *cron.Scheduler) {
	// counting the number of time the job runs and use remove by reference function to remove the job
	if job.RunCount() > 10 {
		s.RemoveByReference(job)
	}

	if len(s.Jobs()) <= 1 {
		// it will unblocks that blocking channel of the instance s
		s.StopBlockingChan()
	}
}

func GetWeatherData(url string, s *cron.Scheduler) {
	var weather Weather

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(respdata, &weather)
	if err != nil {
		panic(err)
	}

	if weather.Name == "" {
		fmt.Println("Not a valid input")
		s.StopBlockingChan()
		return
	}

	fmt.Printf("Weather report for %q at %v\n", weather.Name, time.Now().Format("01-02-2006 03:04:05 PM"))
	fmt.Printf("Current Tempature: %v%c celcius \n", weather.Main.Temp, byte(176))
	fmt.Println("Current Wind Speed:", weather.Wind.Speed)
	fmt.Printf("Current Weather Description: %v \n\n", weather.Weather[0].Description)
}
