package csv

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type CSV struct {
	Records []CSVRow
}

type CSVRow struct {
	Company          string `csv:"Company"`
	Person           string `csv:"Person"`
	Name             string `csv:"Name"`
	DeviceType       string `csv:"Device type"`
	MACAddress       string `csv:"MAC address"`
	Registered       string `csv:"Registered"`
	Status           string `csv:"Status"`
	UUIDCreationDate string `csv:"UUID creation date"`
	DownloadDate     string `csv:"Download date"`
	HotDesking       string `csv:"Hot desking"`
	HotDeskingID     string `csv:"Hot desking ID"`
	HotDeskingPhone  string `csv:"Hot desking phone"`
	Location         string `csv:"Location"`
	Group            string `csv:"Group"`
	Comment          string `csv:"Comment"`
	Firmware         string `csv:"Firmware"`
}

type RowIterator struct {
	csvFile  *CSV
	position int
}

/*
	function for object creation from string

need to add handling of extre double quotes in csv file.
*/
func New(csvString string) (*CSV, error) {
	reader := csv.NewReader(strings.NewReader(csvString))
	reader.TrimLeadingSpace = true
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	// Read the remaining rows
	var csvRows []CSVRow
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("error is coming here inner::")
			fmt.Printf("%v", err)
			return nil, err
		}

		length := len(row)
		fmt.Println(length)

		csvRow := CSVRow{
			Company:          row[0],
			Person:           row[1],
			Name:             row[2],
			DeviceType:       row[3],
			MACAddress:       row[4],
			Registered:       row[5],
			Status:           row[6],
			UUIDCreationDate: row[7],
			DownloadDate:     row[8],
			HotDesking:       row[9],
			HotDeskingID:     row[10],
			HotDeskingPhone:  row[11],
			Location:         row[12],
			Group:            row[13],
			Comment:          row[14],
			Firmware:         row[15],
		}
		csvRows = append(csvRows, csvRow)
	}

	csvObj := &CSV{
		Records: csvRows,
	}

	return csvObj, nil
}

// Implementing CSVContract methods

func (c *CSV) RowIterator(pos int) RowIterator {
	return RowIterator{
		csvFile:  c,
		position: pos,
	}
}

func (c *CSV) Incorporate(other CSV) {
	c.Records = append(c.Records, other.Records...)
}

func (c *CSV) ToStringRFC4180() string {
	var result string

	// Write the header row
	header := []string{
		"Company",
		"Person",
		"Name",
		"Device type",
		"MAC address",
		"Registered",
		"Status",
		"UUID creation date",
		"Download date",
		"Hot desking",
		"Hot desking ID",
		"Hot desking phone",
		"Location",
		"Group",
		"Comment",
		"Firmware",
	}
	result += strings.Join(header, ",") + "\n"

	// Write the data rows
	for _, row := range c.Records {
		data := []string{
			row.Company,
			row.Person,
			row.Name,
			row.DeviceType,
			row.MACAddress,
			row.Registered,
			row.Status,
			row.UUIDCreationDate,
			row.DownloadDate,
			row.HotDesking,
			row.HotDeskingID,
			row.HotDeskingPhone,
			row.Location,
			row.Group,
			row.Comment,
			row.Firmware,
		}
		result += strings.Join(data, ",") + "\n"
	}

	return result
}

// Implementing RowIteratorContract methods

func (it *RowIterator) Get() []string {
	// if it.position >= len(it.csvFile.Records) {
	// 	return nil
	// }

	fmt.Println("================")
	fmt.Println("CSV records len")
	fmt.Println(len(it.csvFile.Records))
	fmt.Println("csv iterator pos")
	fmt.Println(it.position)
	fmt.Println("================")

	if it.position >= len(it.csvFile.Records) {
		return nil
	}
	return []string{
		it.csvFile.Records[it.position].Company,
		it.csvFile.Records[it.position].Person,
		it.csvFile.Records[it.position].Name,
		it.csvFile.Records[it.position].DeviceType,
		it.csvFile.Records[it.position].MACAddress,
		it.csvFile.Records[it.position].Registered,
		it.csvFile.Records[it.position].Status,
		it.csvFile.Records[it.position].UUIDCreationDate,
		it.csvFile.Records[it.position].DownloadDate,
		it.csvFile.Records[it.position].HotDesking,
		it.csvFile.Records[it.position].HotDeskingID,
		it.csvFile.Records[it.position].HotDeskingPhone,
		it.csvFile.Records[it.position].Location,
		it.csvFile.Records[it.position].Group,
		it.csvFile.Records[it.position].Comment,
		it.csvFile.Records[it.position].Firmware,
	}
}

func (it *RowIterator) Next() bool {
	it.position++
	fmt.Println(it.position)
	return it.position < len(it.csvFile.Records)
}
func (it *RowIterator) Pos() int {
	return it.position
}
