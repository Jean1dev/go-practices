package entity

import (
	"testing"
)

func Test_If_I_Get_An_Error_If_Id_IsBlank(t *testing.T) {
	order := Order{}
	err := order.Validate()

	if err == nil {
		t.Error("Expected error, got nil")
	}

	t.Logf("Okkkk")
}
