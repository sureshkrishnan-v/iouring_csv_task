package csvfile

import (
	"fmt"
	"os"
)

func OpenCsvFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return file, fmt.Errorf("error opening csv file")
	}
	return file, nil
}
