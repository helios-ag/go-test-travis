package go_test_travis

import "fmt"

func triggerError(showError bool) (string, error) {

	if showError {
		return "", fmt.Errorf("error triggered")
	}

	result := "All is OK"

	return result, nil
}
