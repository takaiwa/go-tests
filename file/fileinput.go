package file

import (
	"encoding/csv"
	"io"
	"os"
)

type Data []string

func LoadCsv(path string) (map[string]Data, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	maps, err := csvToMap(file)
	return maps, err
}

func csvToMap(in io.Reader) (map[string]Data, error) {
	r := csv.NewReader(in)
	table := map[string]Data{}
	var keys []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if keys == nil {
			keys = record
		} else {
			for index, k := range keys {
				d, exists := table[k]
				data := Data{}
				if exists {
					data = append(d, record[index])
				} else {
					data = append(data, record[index])
				}
				table[k] = data
			}
		}
	}

	return table, nil
}
