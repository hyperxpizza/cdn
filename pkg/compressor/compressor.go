package compressor

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func CompressFile(data []byte) (int, []byte, error) {
	var buffer bytes.Buffer
	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestSpeed)
	if err != nil {
		return 0, nil, err
	}
	size, err := writer.Write(data)
	if err != nil {
		return 0, nil, err
	}

	err = writer.Flush()
	if err != nil {
		return 0, nil, err
	}

	err = writer.Close()
	if err != nil {
		return 0, nil, err
	}

	return size, buffer.Bytes(), nil
}

func DecompressFile(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}
