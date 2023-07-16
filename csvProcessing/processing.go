package csvprocessing

import (
	"fmt"
	"maccsv/csv"
	"strings"
	"time"
)

func GetNewPhonesRegistered(iterator csv.RowIterator, lastDownloadTime time.Time) []string {

	columnNames := iterator.Get()

	registeredColumnIndex := -1
	for i, columnName := range columnNames {
		if columnName == "UUID creation date" {
			registeredColumnIndex = i
			break
		}
	}

	if registeredColumnIndex == -1 {
		fmt.Println("Error: 'UUID creation date' column not found")
		return nil
	}

	newPhonesRegistered := []string{}

	for iterator.Next() {
		row := iterator.Get()
		if len(row) > registeredColumnIndex && row[registeredColumnIndex] != "" {

			var registrationTime time.Time
			var err error

			input := row[registeredColumnIndex]

			monthMap := map[string]string{
				"Sept": "Sep",
			}

			for k, v := range monthMap {
				input = strings.Replace(input, k, v, -1)
			}

			registrationTime, err = time.Parse("2 Jan 2006 15:04:05", input)

			if err != nil {
				fmt.Println("Error parsing date:", err)
			}

			fmt.Println("Last Download time:", lastDownloadTime)

			if registrationTime.After(lastDownloadTime) {
				newPhonesRegistered = append(newPhonesRegistered, row[registeredColumnIndex])
			}
		}
	}

	return newPhonesRegistered
}
