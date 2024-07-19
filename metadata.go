package business_errors

import (
	"fmt"
	"github.com/ValGoldun/logger"
	"strings"
)

type Metadata map[string]string

func (m Metadata) LoggerFields() logger.Fields {
	var fields = make(logger.Fields, len(m))

	var i int
	for key, value := range m {
		fields[i] = logger.Field{
			Key:   key,
			Value: value,
		}

		i++
	}

	return fields
}

func (m Metadata) String() string {
	var fields = make([]string, len(m))

	var i int
	for key, value := range m {
		fields[i] = fmt.Sprintf("%s: %s", key, value)

		i++
	}

	return strings.Join(fields, ", ")
}
