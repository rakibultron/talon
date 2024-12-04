package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// ConvertUUIDToPGXUUID converts a uuid.UUID to pgtype.UUID
func ConvertUUIDToPGXUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id, // Direct assignment since both are [16]byte
		Valid: true,
	}
}
