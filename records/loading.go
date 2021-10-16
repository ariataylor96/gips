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

func getXBytes(buf *bufio.Reader, n int) (data []byte, rerr error) {
	data = make([]byte, n)

	read, err := io.ReadFull(buf, data)
	if err != nil {
		rerr = err
		return
	}

	if read != n {
		panic(errors.New(fmt.Sprintf("Read fewer bytes than expected: %v != %v", read, n)))
	}

	return
}

func parseXBytes(buf *bufio.Reader, n int) (res uint32, rerr error) {
	parsed, err := getXBytes(buf, n)
	if err != nil {
		rerr = err
		return
	}

	for len(parsed) < 4 {
		parsed = append(parsed, 0x0)
	}

	res = binary.BigEndian.Uint32(parsed)
	return
}

func atEOF(buf *bufio.Reader) bool {
	eofSignature := []byte{'E', 'O', 'F'}
	next3Bytes, err := buf.Peek(3)
	if err != nil {
		if err == io.EOF {
			return true
		}
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

	header, _ := getXBytes(buf, 5)
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
	offset, err := parseXBytes(buf, 3)
	if err != nil {
		panic(err)
	}
	res.Offset = offset

	size, err := parseXBytes(buf, 2)
	if err != nil {
		panic(err)
	}
	res.Size = uint16(size)

	if res.Size == 0 {
		res.IsRLE = true

		rleSize, _ := parseXBytes(buf, 2)
		res.RLESize = uint16(rleSize)

		rawVal, _ := getXBytes(buf, 1)
		res.rawValue = rawVal
	} else {
		rawVal, _ := getXBytes(buf, int(res.Size))
		res.rawValue = rawVal
	}

	return
}
