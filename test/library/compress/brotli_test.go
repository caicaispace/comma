package compress_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/andybalholm/brotli"
)

func TestDecode(t *testing.T) {
	str := "1121212,43363565,24234354,2324324,1224232,123125354,9899797,9899797,9899797,9899797,9899797,9899797"
	content := bytes.Repeat([]byte(str), 1)
	encoded, _ := Encode(content, brotli.WriterOptions{Quality: 2})
	decoded, err := Decode(encoded)
	if err != nil {
		t.Errorf("Decode: %v", err)
	}
	t.Log(len(encoded))
	t.Log(len(decoded))
	t.Log(float32(len(encoded) / len(decoded)))
}

// Encode returns content encoded with Brotli.

func Encode(content []byte, options brotli.WriterOptions) ([]byte, error) {
	var buf bytes.Buffer
	writer := brotli.NewWriterOptions(&buf, options)
	_, err := writer.Write(content)
	if closeErr := writer.Close(); err == nil {
		err = closeErr
	}
	return buf.Bytes(), err
}

// Decode decodes Brotli encoded data.
func Decode(encodedData []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(encodedData))
	return ioutil.ReadAll(r)
}
