package service

import (
	// Standard Libs
	"fmt"
	"os"
	"time"

	// Third Party Libs
	cron "github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

var (
	apiKey string
	s      *cron.Scheduler
	url    string
)

func init() {
	godotenv.Load(".env")
	apiKey = os.Getenv("API_KEY")

	url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=metric&", apiKey)
	s = cron.NewScheduler(time.Local)
}
