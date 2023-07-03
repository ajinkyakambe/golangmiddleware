package csvprocessing

import (
	"fmt"
	"maccsv/csv"
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

	fmt.Printf("Register column index %v", registeredColumnIndex)

	newPhonesRegistered := []string{}

	for iterator.Next() {
		row := iterator.Get()
		if len(row) > registeredColumnIndex && row[registeredColumnIndex] != "" {
			layout := []string{
				"2 Jan 2006 15:04:05",
				"1 Sept 2020 17:08:33",
				"18 Jun 2020 12:22:25",
				"2 Jan 2006 15:04:05 MST",
			}

			var registrationTime time.Time
			var err error

			fmt.Println(row[registeredColumnIndex])
			for _, l := range layout {
				registrationTime, err = time.Parse(l, row[registeredColumnIndex])
				if err == nil {
					break // Exit the loop if parsing is successful
				}
			}

			if err != nil {
				fmt.Println("Error parsing date:", err)

			}

			if registrationTime.After(lastDownloadTime) {
				newPhonesRegistered = append(newPhonesRegistered, row[registeredColumnIndex])
			}
		}
	}

	return newPhonesRegistered
}
