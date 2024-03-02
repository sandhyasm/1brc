package brc

import "io"

// Paris=1.2/20.6/12.2

func Process(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
