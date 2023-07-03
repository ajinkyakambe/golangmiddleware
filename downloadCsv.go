package main

import (
	csvOriginal "encoding/csv"
	"fmt"
	"io"
	"maccsv/csv"
	csvProcessing "maccsv/csvProcessing"
	util "maccsv/etc"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var lastDownloadTime time.Time

func autoDownloadCSV() {

	downloadTime := time.Now().Add(5 * time.Second)
	// Currently for testing it is hard coded
	// after that we will take it from config file.

	var wg sync.WaitGroup

	multiServer := util.ReadMultiServerConfig()

	for serverNumber, url := range *multiServer {

		wg.Add(1)
		go func(u string, serverNumber int) {
			defer wg.Done()

			time.Sleep(downloadTime.Sub(time.Now()))

			currentTime := time.Now()
			formattedTime := currentTime.Format("02_01_2006")

			filename := fmt.Sprintf("downloaded_server%d_%v.csv", serverNumber, formattedTime)

			err := DownloadCSV(u, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", u, err)
			}

			fmt.Println(filename)

			//csvString
			csvString, err := LoadCSVAsString(filename)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println(csvString)

			// Create a new CSV object
			csvObj, err := csv.New(csvString)
			if err != nil {
				fmt.Println("Error is it:", err)
				return
			}

			iterator := csvObj.RowIterator(0)

			newPhones := csvProcessing.GetNewPhonesRegistered(iterator, lastDownloadTime)
			fmt.Println(newPhones)

		}(url, serverNumber)
	}

	wg.Wait()

}

func DownloadCSV(url, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s\n", filename)
	layout := "02-01-2006 15:04:05"
	currentTime := time.Now()

	formattedTimeStr := currentTime.Format(layout)
	formattedTime, err := time.Parse(layout, formattedTimeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return err
	}

	lastDownloadTime = formattedTime
	return nil
}

func LoadCSVAsString(filename string) (string, error) {
	// Open the CSV file.
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a reader for the CSV file.
	reader := csvOriginal.NewReader(file)

	// Read the CSV file into a string.
	var csvData string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Remove newlines from quoted fields
		for i, field := range line {
			field = strings.ReplaceAll(field, "\n", " ")
			field = strings.ReplaceAll(field, ",", " /")
			line[i] = field
		}

		csvData += fmt.Sprintf("%s\n", strings.Join(line, ","))
	}

	return csvData, nil
}
