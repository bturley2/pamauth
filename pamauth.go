package pamauth

import (
	"fmt"

	"github.com/msteinert/pam"
)

func PerformPamAuth() error {
	// initiate PAM transaction
	pam.StartFunc("", "", func(s pam.Style, msg string) (string, error) {
		return "", nil
	})

	//
	fmt.Println("It Worked!")
	return nil
}
