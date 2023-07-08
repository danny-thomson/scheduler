package service

import (
	// Standard Libs
	"fmt"

	// Third Party Libs
	"cron/service/weather"
)

func GetWeatherByCity(city string) {
	// starting the job
	locationJob, _ := s.Every(2).Second().Do(func() {
		url := fmt.Sprintf("%s&q=%s", url, city)
		weather.GetWeatherData(url, s)
	})

	s.Every(2).Second().Do(func() { weather.Stop(locationJob, s) })

	s.StartBlocking()
}

func GetWeatherByPincode(pincode string) {
	// it will execute this task for 10 times only
	s.Every(2).Second().LimitRunsTo(10).Do(func() {
		url := fmt.Sprintf("%s&zip=%s,in", url, pincode)
		weather.GetWeatherData(url, s)
	})

	//LimitRunsTo function will only remove the job after it excutes the jobs but only stopblockingchan func will unblocks the channel
	s.Every(2).Second().Do(func() {
		if s.Len() <= 1 {
			s.StopBlockingChan()
		}
	})

	s.StartBlocking()
}
