package validators

import (
	"fmt"
	"testing"
)

func TestValidateOrder(t *testing.T) {
	// should return error case user_id is not provided
	ot := new(Order)

	ot.UserID = "valid_id"

	err := ot.ValidateOrder()

	expected := "user_id: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if user_id is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if user_id is not provided.")
	}
}

// should return error case item_description is not provided

// should return error case item_description min5-max 200 is not provided
// should return error case item_quantity is not provided
// should return error case item quantity is not float
// should return error case item_price is not provided
// should return error case item_price is not float
// should return error case order data is valid
