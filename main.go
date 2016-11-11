package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nomkhonwaan/logger/config"
	"github.com/spf13/viper"
)

var (
	lg *log.Logger
	f  *os.File
)

type target struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func init() {
	f, err := os.OpenFile(
		time.Now().Format(viper.GetString("output_logs.path")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	lg = log.New(os.Stdout, "", 0)
	lg.SetOutput(f)
}

func main() {
	defer f.Close()

	var targets []target
	err := viper.UnmarshalKey("targets", &targets)
	if err != nil {
		lg.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
		triggerCall(targets)
	}
}

func triggerCall(in []target) {
	done := make(chan struct{})

	for _, n := range in {
		go func(in2 target) {
			call(in2)
			done <- struct{}{}
		}(n)
	}
	for j := 0; j < len(in); j++ {
		<-done
	}
}

func call(in target) {
	max := viper.GetInt("calls_per_second")
	done := make(chan struct{})

	for i := 0; i < max; i++ {
		go func(in2 target) {
			start := time.Now()

			client := newClient()
			req, _ := http.NewRequest(in2.Method, in2.URL, nil)
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("[ERROR] an error has occurred: %s", err)
			}
			defer resp.Body.Close()

			stop := time.Now()
			elapsed := stop.Sub(start)

			w(map[string]interface{}{
				"URL":           in2.URL,
				"method":        in2.Method,
				"error-message": err,
				"requested-at":  start,
				"received-at":   stop,
				"elasped-ns":    elapsed.Nanoseconds(),
				"elapsed-s":     elapsed.Seconds(),
				"elapsed-ms":    (elapsed.Nanoseconds() % 1e9) / 1e6,
				"elapsed-m":     elapsed.Minutes(),
			})

			if config.LoggerConfig.Debug {
				log.Printf("[INFO] calling [%s:%s] at [%s], response in [%.2fs]", in2.Method, in2.URL, start.Format(time.RFC3339), elapsed.Seconds())
			}

			done <- struct{}{}
		}(in)
	}
	for j := 0; j < max; j++ {
		<-done
	}
}

func w(in interface{}) error {
	if viper.GetString("output_logs.type") == "JSON" {
		var b []byte
		var err error

		if _, ok := in.(target); ok {
			b, err = json.Marshal(in.(target))
		} else {
			b, err = json.Marshal(in)
		}

		if err != nil {
			return err
		}
		lg.Println(string(b))
	} else {
		lg.Println(in)
	}
	return nil
}

func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
