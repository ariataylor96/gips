package ips

import "errors"

var FILE_HEADER []byte = []byte{'P', 'A', 'T', 'C', 'H'}

func ValidateHeader(data *[]byte) error {
	data_header := (*data)[0:5]

	for idx, val := range data_header {
		if val != FILE_HEADER[idx] {
			return errors.New("IPS file does not have a valid header")
		}
	}

	return nil
}
