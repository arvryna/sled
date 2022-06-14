package utils

import "os"

func GetContentsOfFile(filePath string) string {
	options, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(options)
}

// https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permissions
func CreateFolderIfNotExist(path string) error {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0700)
	}
	return err
}
