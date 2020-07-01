package pamauth

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/msteinert/pam"
)

func PerformPamAuth() error {
	// initiate PAM transaction
	// prompts directly from terminal this program is run from
	// TODO: redirect prompts to user's client using channel
	pam.StartFunc("", "", func(s pam.Style, msg string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			return "", nil
		case pam.PromptEchoOn:
			fmt.Print(msg + " ")
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return "", err
			}
			return input[:len(input)-1], nil
		case pam.ErrorMsg:
			log.Print(msg)
			return "", nil
		case pam.TextInfo:
			fmt.Println(msg)
			return "", nil
		}

		return "", errors.New("Unrecognized message style")
	})

	//
	fmt.Println("It Worked!")
	return nil
}
