package csvfile

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/iouring_task/models"
)

func WriteCsvFile(data []models.RelianceData) error {
	file, err := os.Create("output.xlsx")
	if err != nil {
		return fmt.Errorf("error writing csv%v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Date", "Open", "High", "Low", "Close", "AdjClose", "Volume",
		"SMAOpen", "SMAHigh", "SMALow", "SMAClose", "SMAAdjClose", "SMAVolumeI",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing csv%v", err)
	}

	// Writing CSV Data
	for _, record := range data {
		row := []string{
			record.Date,
			strconv.FormatFloat(record.Open, 'f', 2, 64),
			strconv.FormatFloat(record.High, 'f', 2, 64),
			strconv.FormatFloat(record.Low, 'f', 2, 64),
			strconv.FormatFloat(record.Close, 'f', 2, 64),
			strconv.FormatFloat(record.AdjClose, 'f', 2, 64),
			strconv.FormatFloat(record.Volume, 'f', 2, 64),
			strconv.FormatFloat(record.SMAOpen, 'f', 2, 64),
			strconv.FormatFloat(record.SMAHigh, 'f', 2, 64),
			strconv.FormatFloat(record.SMALow, 'f', 2, 64),
			strconv.FormatFloat(record.SMAClose, 'f', 2, 64),
			strconv.FormatFloat(record.SMAAdjClose, 'f', 2, 64),
			strconv.FormatFloat(record.SMAVolumeI, 'f', 2, 64),
		}
		if err := writer.Write(row); err != nil {
			fmt.Errorf("error writing csv%v", err)
		}
	}
	return nil
}
