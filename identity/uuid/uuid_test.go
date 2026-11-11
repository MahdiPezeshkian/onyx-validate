package uuid_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/identity/uuid"
)

func TestUUID(t *testing.T) {
	v := "550e8400-e29b-41d4-a716-446655440000"
	if err := uuid.Validate(v); err != nil {
		t.Fatal(err)
	}
	ver, err := uuid.Version(v)
	if err != nil || ver != 4 {
		t.Fatalf("version=%d %v", ver, err)
	}
	if uuid.IsValid("550e8400e29b41d4a716446655440000") {
		t.Fatal("compact form not accepted in this version")
	}
}
