package main

import (
	"os"
	"secret/functions"
)

func main() {
	command, file_name := functions.Command()

	file_content, file_read_err := os.ReadFile(file_name)

	if file_read_err != nil {
		panic(file_read_err)
	}

	switch command {
	case "hash":
		hashed_content := functions.Hash_string(string(file_content))
		functions.Trunc_write(file_name, hashed_content)

	case "decode":
		file_content, file_read_err := os.ReadFile(file_name)

		if file_read_err != nil {
			panic(file_read_err)
		}

		decoded_content, decoded_err := functions.Decode_string(string(file_content))

		if decoded_err != nil {
			panic(decoded_err)
		}

		functions.Trunc_write(file_name, decoded_content)

	}

}
