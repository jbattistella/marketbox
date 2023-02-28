package makeCSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ListHeaders(t string) {
	csvFile, _ := os.Open(t)
	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	line, error := reader.Read()
	if error == io.EOF {
		fmt.Println("EOF")
	} else if error != nil {
		log.Fatal(error)
	}
	for index, x := range line {
		fmt.Printf("%d:%s\n", index, x)
	}

}
