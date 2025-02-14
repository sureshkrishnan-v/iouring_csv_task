package csvfile

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(file *os.File) ([][]string , error){
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error Occured When Reading File", err.Error())
	}
	return records , err
}
