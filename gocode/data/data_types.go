package data

import (
	"github.com/google/uuid"
    "time"
)

type EntityType struct {
	Id uuid.UUID
	Name string
	Description string
}

type Entity struct {
	Id uuid.UUID
	CreatorId uuid.UUID
	Name string
	Type uuid.UUID
	Details string
}

type User struct {
	Id uuid.UUID
	Name string
}

type Connection struct {
	Id uuid.UUID
	CreatorId uuid.UUID
	EntityA uuid.UUID
	EntityB uuid.UUID
	ConnectingEntity uuid.UUID
	StartDate time.Time
	EndDate time.Time
	Details string
}

type DuplicateEntities struct {
    Id uuid.UUID
	CreatorId uuid.UUID
	EntityId uuid.UUID
	DuplicateEntityId uuid.UUID
}