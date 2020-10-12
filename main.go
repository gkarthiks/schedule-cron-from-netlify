package main

import (
	"k8s.io/klog"
	"os"
	"strconv"
	"time"
)

var (
	frequency int
	err error
	count = 0
)
const (
	DefaultScheduleFrequencyInMinutes = 15
)
func init() {
	frequencyStr, avail := os.LookupEnv("SCHEDULE_FREQ")
	if !avail {
		klog.Infof("no schedule is provided, defaulting with %d minutes", DefaultScheduleFrequencyInMinutes)
		frequency = DefaultScheduleFrequencyInMinutes
	} else {
		frequency, err = strconv.Atoi(frequencyStr)
		if err!= nil {
			klog.Errorf("error while parsing string to int")
		}
	}
}

func main() {
	// create a routine
	done := make(chan bool)
	go runOnSchedule(frequency)
	<- done
}

func runOnSchedule(frequency int) {
	for ;; {
		count = count+1
		time.Sleep(time.Duration(frequency) * time.Second)
		klog.Infof("Ran for %d time @ %s", count, time.Now().Format(time.RFC3339))
	}
}
