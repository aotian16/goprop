// prop.go
package goprop

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Properties map[string]string

var NotExistKeyError error = errors.New("NotExistKeyError")

func NewProp() Properties {
	return make(Properties)
}

// get value as string.
// different to prop[key],the return string value is Unquoted.
func (this *Properties) Get(key string) (string, error) {
	v, ok := (*this)[key]

	if !ok {
		return "", NotExistKeyError
	}

	return v, nil
}

// get value as int
func (this *Properties) Atoi(key string) (int, error) {
	v, ok := (*this)[key]

	if !ok {
		return 0, NotExistKeyError
	}

	return strconv.Atoi(v)
}

// get value as bool
func (this *Properties) ParseBool(key string) (bool, error) {
	v, ok := (*this)[key]

	if !ok {
		return false, NotExistKeyError
	}

	return strconv.ParseBool(v)
}

// get value as float
func (this *Properties) ParseFloat(key string) (float64, error) {
	v, ok := (*this)[key]

	if !ok {
		return 0, NotExistKeyError
	}

	return strconv.ParseFloat(v, 64)
}

// get value as int
func (this *Properties) ParseInt(key string) (int64, error) {
	v, ok := (*this)[key]

	if !ok {
		return 0, NotExistKeyError
	}

	return strconv.ParseInt(v, 10, 64)
}

// get value as uint
func (this *Properties) ParseUint(key string) (uint64, error) {
	v, ok := (*this)[key]

	if !ok {
		return 0, NotExistKeyError
	}

	return strconv.ParseUint(v, 10, 64)
}

// load properties from file
func LoadFile(fileName string) (Properties, error) {
	p := make(Properties)
	file, err := os.Open(fileName)

	if err != nil {
		return p, err
	}

	defer file.Close()

	return Load(file)
}

// load properties from reader
func Load(reader io.Reader) (Properties, error) {
	p := make(Properties)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		kv := strings.SplitN(scanner.Text(), "=", 2)

		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			p[k] = v
		}
	}

	if scanner.Err() != nil {
		return p, scanner.Err()
	}

	return p, nil
}

// save properties to file
// * file must not exist, because it's dangerous to overwrite
func SaveFile(prop Properties, fileName string) error {
	file, errOpen := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL, os.ModePerm)

	if errOpen != nil {
		return errOpen
	}

	defer file.Close()

	return Save(prop, file)
}

// save properties to writer
func Save(prop Properties, writer io.Writer) error {
	bufWriter := bufio.NewWriter(writer)

	for k, v := range prop {
		_, err := bufWriter.WriteString(k + "=" + v + "\n")

		if err != nil {
			return err
		}
	}

	errFlush := bufWriter.Flush()
	if errFlush != nil {
		return errFlush
	}

	return nil
}
