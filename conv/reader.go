package conv

import (
	"encoding/xml"
	"io"
)

func Read(r io.Reader) (*Scheme, error) {

	var scheme Scheme
	if err := xml.NewDecoder(r).Decode(&scheme); err != nil {
		return nil, err
	}

	return &scheme, nil
}
