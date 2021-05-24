package storage

import "ports/internal/pkg/pb"

type Ports interface {
	Get(code string) pb.Port
	List(code ...string) []pb.Port
	Delete(code string) error
	Save(port pb.Port)
}
