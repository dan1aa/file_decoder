package functions

import "encoding/base64"

func Hash_string(input string) []byte {
	hashed_string := base64.StdEncoding.EncodeToString([]byte(input))
	return []byte(hashed_string)
}

func Decode_string(hashed_string string) ([]byte, error) {
	decoded_bytes, err := base64.StdEncoding.DecodeString(hashed_string)
	if err != nil {
		return []byte{}, err
	}
	return decoded_bytes, nil
}
