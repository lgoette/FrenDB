package data
import (
	"github.com/google/uuid"
)

type Entity struct {
	Id uuid.UUID
	CreatorId uuid.UUID
	Name string
	Type string
	Details string
}