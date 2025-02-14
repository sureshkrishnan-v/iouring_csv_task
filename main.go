package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iouring_task/models"
	csvfile "github.com/iouring_task/utils/FileOperation"
)

func main() {
	fmt.Println("Started")

	file, err := os.Open("RELIANCE.csv")
	if err != nil {
		fmt.Println("Error Reading File:", err.Error())
		return
	}
	defer file.Close()
	records, err := csvfile.ReadCsvFile(file)
	if err != nil {
		fmt.Println("Error reading CSV File", err.Error())
	}
	relinceDataArr, err := PopulateDataFromCSV(records)

	if err != nil {
		fmt.Println("Error in Reliance Data Population")
	}

	FinalData := CalculateSMA(relinceDataArr)
	err = csvfile.WriteCsvFile(FinalData)
	if err != nil {
		fmt.Println(err)
	}
}

func CalculateSMA(relianceDataArr []models.RelianceData) []models.RelianceData {

	for i, _ := range relianceDataArr {
		relianceDataArr[i].SMAOpen = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.Open }, i, 10)
		relianceDataArr[i].SMALow = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.Low }, i, 10)
		relianceDataArr[i].SMAClose = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.Close }, i, 10)
		relianceDataArr[i].SMAAdjClose = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.AdjClose }, i, 10)
		relianceDataArr[i].SMAVolumeI = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.Volume }, i, 10)
		relianceDataArr[i].SMAHigh = calculateSMAForField(relianceDataArr, func(d models.RelianceData) float64 { return d.High }, i, 10)

	}

	return relianceDataArr
}

func calculateSMAForField(relianceDataArr []models.RelianceData, fieldSelector func(models.RelianceData) float64, index int, period int) float64 {
	if index < period-1 {
		return 0.0
	}

	sum := 0.0
	for i := index - period + 1; i <= index; i++ {
		sum += fieldSelector(relianceDataArr[i])
	}

	return sum / float64(period)
}

func PopulateDataFromCSV(records [][]string) ([]models.RelianceData, error) {
	var relinceDataArr []models.RelianceData
	for i := 1; i < len(records); i++ {
		var relianceData models.RelianceData
		for j := 0; j < len(records[i]); j++ {
			relianceData.Date = records[i][0]
			relianceData.Open = ParseFloat(records[i][1])
			relianceData.High = ParseFloat(records[i][2])
			relianceData.Low = ParseFloat(records[i][3])
			relianceData.Close = ParseFloat(records[i][4])
			relianceData.AdjClose = ParseFloat(records[i][5])
			relianceData.Volume = ParseFloat(records[i][6])
			relianceData.SMAOpen = ParseFloat(records[i][7])
			relianceData.SMAHigh = ParseFloat(records[i][8])
			relianceData.SMALow = ParseFloat(records[i][9])
			relianceData.SMAClose = ParseFloat(records[i][10])
			relianceData.SMAAdjClose = ParseFloat(records[i][11])
			relianceData.SMAVolumeI = ParseFloat(records[i][12])
		}
		relinceDataArr = append(relinceDataArr, relianceData)

	}
	return relinceDataArr, nil
}

func ParseFloat(record string) float64 {
	var val float64
	var err error
	if record != "" {
		val, err = strconv.ParseFloat(record, 64)
		if err != nil {
			fmt.Println("error parsing float", err.Error())
		}
	}
	return val
}
