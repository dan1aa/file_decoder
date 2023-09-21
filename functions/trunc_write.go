package functions

import "os"

func Trunc_write(file_name string, file_content []byte) {
	current_file, current_file_err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if current_file_err != nil {
		panic(current_file_err)
	}

	trunc_err := current_file.Truncate(0)

	if trunc_err != nil {
		panic(trunc_err)
	}

	write_hash_err := os.WriteFile(file_name, file_content, 0644)

	if write_hash_err != nil {
		panic(write_hash_err)
	}
}
