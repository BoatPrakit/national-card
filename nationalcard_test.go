package nationalcard

import "testing"

func TestValidatePasswordError(t *testing.T) {
	testcases := []struct {
		given    string
		expected string
	}{
		{
			given:    "123456789012",
			expected: "id digits incorrect",
		},
		{
			given:    "1234567890121",
			expected: "",
		},
		{
			given:    "1234567890122",
			expected: "id incorrect",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidateThaiID(testcase.given)

			if actual == nil {
				println("passed")
				return
			}

			if testcase.expected != actual.Error() {
				t.Errorf("given a password %s expected error is %v but actual error is %v", testcase.given, testcase.expected, actual)
			}
		})
	}
}

// func TestValidatePasswordValid(t *testing.T) {
// 	testcases := []struct {
// 		given string
// 	}{
// 		{
// 			given: "@bcdeF12",
// 		},
// 	}

// 	for _, testcase := range testcases {
// 		t.Run("", func(t *testing.T) {
// 			actual := ValidatePassword(testcase.given)

// 			if actual != nil {
// 				t.Errorf("given a password %s should valid but actual error is %v", testcase.given, actual)
// 			}
// 		})
// 	}
// }
