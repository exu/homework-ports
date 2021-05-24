package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"ports/internal/pkg/pb"
)

type Callback func(port *pb.Port) error

func ProcessPortsJSON(input io.Reader, callback Callback) error {
	dec := json.NewDecoder(input)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return fmt.Errorf("decoding token error: %w", err)
	}

	// parse map key -> .Token, value -> .Decode
	for dec.More() {
		t, err := dec.Token()
		log.Printf("Scanning %s\n", t)
		if err != nil {
			return fmt.Errorf("decode token error: %w", err)
		}
		var port pb.Port
		err = dec.Decode(&port)
		if err != nil {
			return fmt.Errorf("decoding error: %w", err)
		}

		fmt.Printf("PORT: %+v\n", &port)

		err = callback(&port)
		if err != nil {
			return fmt.Errorf("calling callback: %w", err)
		}
	}

	// read closing bracket
	_, err = dec.Token()
	return err
}
