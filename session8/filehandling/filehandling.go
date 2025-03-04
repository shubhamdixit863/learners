package filehandling

import "os"

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil

}

//func WriteFunction(filename string) (string, error) {
//	data, err := os.WriteFile()
//	if err != nil {
//		return "", err
//	}
//
//	return string(data), nil
//
//}
