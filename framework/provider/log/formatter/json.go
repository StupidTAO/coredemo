package formatter

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/gohade/hade/framework/contract"
	"time"
)

func JsonFormatter(level contract.LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte(msg))
	bf.Write([]byte{':'})
	c, err := json.Marshal(fields)
	if err != nil {
		return bf.Bytes(), errors.Wrap(err, "json format error")
	}

	bf.Write(c)
	return bf.Bytes(), nil
}
