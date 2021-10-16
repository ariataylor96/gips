package records

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"gips/validators/ips"
	"io"
	"os"
)

func getXBytes(buf *bufio.Reader, n int) []byte {
	data := make([]byte, n)

	read, err := io.ReadFull(buf, data)
	if err != nil {
		panic(err)
	}

	if read != n {
		panic(errors.New(fmt.Sprintf("Read fewer bytes than expected: %v != %v", read, n)))
	}

	return data
}

func parseXBytes(buf *bufio.Reader, n int) uint32 {
	parsed := getXBytes(buf, n)

	for len(parsed) < 4 {
		parsed = append(parsed, 0x0)
	}

	return uint32(binary.BigEndian.Uint16(parsed))
}

func atEOF(buf *bufio.Reader) bool {
	eofSignature := []byte{'E', 'O', 'F'}
	next3Bytes, err := buf.Peek(3)
	if err != nil {
		panic(err)
	}

	for idx, val := range next3Bytes {
		if val != eofSignature[idx] {
			return false
		}
	}

	return true
}

func FromFile(fileName string) (res []Record) {
	handle, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	buf := bufio.NewReader(handle)

	header := getXBytes(buf, 5)
	err = ips.ValidateHeader(&header)
	if err != nil {
		panic(err)
	}

	for !atEOF(buf) {
		res = append(res, parseRecord(buf))
	}

	return
}

func parseRecord(buf *bufio.Reader) (res Record) {
	res.Offset = parseXBytes(buf, 3)
	res.Size = uint16(parseXBytes(buf, 2))

	if res.Size == 0 {
		res.IsRLE = true

		res.RLESize = uint16(parseXBytes(buf, 2))
		res.rawValue = getXBytes(buf, 1)
	} else {
		res.rawValue = getXBytes(buf, int(res.Size))
	}

	return
}
