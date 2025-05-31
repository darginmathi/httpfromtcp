package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	file, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	requestLine, err := ParseRequestLine(file)
	if err != nil {
		return nil, err
	}

	return &Request{RequestLine: *requestLine}, nil
}

func ParseRequestLine(data []byte) (*RequestLine, error) {
	var method, requestTarget, httpVersion string
	splitRequest := strings.Split(string(data), "\r\n")
	requestLine := splitRequest[0]
	splitRequestLine := strings.Split(requestLine, " ")

	if len(splitRequestLine) != 3 {
		return nil, fmt.Errorf("invalid number of parts in request line")
	}
	if splitRequestLine[0] == strings.ToUpper(splitRequestLine[0]) {
		method = splitRequestLine[0]
	} else {
		return nil, fmt.Errorf("method not found")
	}
	requestTarget = splitRequestLine[1]
	if splitRequestLine[2] == "HTTP/1.1" {
		httpVersion = "1.1"
	} else {
		return nil, fmt.Errorf("incompatable HTTP version %v", splitRequestLine[2])
	}

	return &RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   httpVersion,
	}, nil
}
