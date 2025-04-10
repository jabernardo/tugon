package core

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

type LogMessage struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"msg"`
	Key     string `json:"k"`
}

func TestLogger(t *testing.T) {
	t.Run("Should log", func(t *testing.T) {
		originalStdout := os.Stdout

		defer func() { os.Stdout = originalStdout }()

		r, w, _ := os.Pipe()
		os.Stdout = w

		Logger().Info("Example", "k", "v")

		w.Close()

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)

		os.Stdout = originalStdout

		var log LogMessage
		err := json.Unmarshal(buf.Bytes(), &log)

		if err != nil {
			t.Error("[core.logger] could not read logs", "err", err)
		}

		if log.Time == "" {
			t.Error("[core.logger] Time doesn't exists in log")
		}

		if log.Message == "" {
			t.Error("[core.logger] Message doesn't exists in log")
		}

		if log.Level != "INFO" {
			t.Errorf("[core.logger] Level expected to be `INFO`, got `%s`", log.Level)
		}

		if log.Key != "v" {
			t.Errorf("[core.logger] `k` expected to have value of `v`, got `%s`", log.Key)
		}
	})
}
