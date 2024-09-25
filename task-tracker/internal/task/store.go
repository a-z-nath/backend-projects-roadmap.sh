package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename string = "tasks.json"





func ReadFromFile() (*List, error){
	// data, err := os.ReadFile(filename)
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("There is no file...\nNew file is creating:\n")
			file, err := os.Create(filename)
			if err != nil {
				fmt.Printf("Error to create file: %v\n", err)
				return nil, err
			}
			defer file.Close()
			fmt.Printf("File has been created.\n\n")

		} else {
			fmt.Printf("Something went wrong while reading file : %v\n\n", err)
			return nil, err
		}
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n\n", err)
		return nil, err
	}
	if len(data) == 0 {
		return &List{}, nil
	}
	tasks := &List{}
	json.Unmarshal(data, tasks)
	return tasks, nil
}

func WriteToFile(content []byte) error {
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		fmt.Printf("Error writing to file: %v\n\n", err)
		return err
	}
	return nil
}