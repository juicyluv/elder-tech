package auth_test

import (
	"context"
	"testing"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/domain"
)

func TestAuthContext(t *testing.T) {
	v := domain.AuthContext{
		ID:   12,
		Type: 2,
		Name: "hello",
	}

	ctx := auth.SetAuthContext(context.Background(), &v)

	data, err := auth.GetAuthContextFromContext(ctx)
	if err != nil {
		t.Fatalf("getting context: %v", err)
	}

	if v != *data {
		t.Fatalf("Want: %#v, got: %#v\n", v, data)
	}
}
