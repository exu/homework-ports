package parser

import (
	"ports/internal/pkg/pb"
	"ports/internal/pkg/test/assert"
	"strings"
	"testing"
)

const jsonStream = `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }

}
`

func TestProcessPortsJSON(t *testing.T) {

	t.Run("should process map based json as stream", func(t *testing.T) {

		input := strings.NewReader(jsonStream)

		ports := []*pb.Port{}
		ProcessPortsJSON(input, func(port *pb.Port) error {
			ports = append(ports, port)
			return nil
		})

		assert.Equals(t, "Ajman", ports[0].Name)
		assert.Equals(t, "Abu Dhabi", ports[1].Name)
	})

}
