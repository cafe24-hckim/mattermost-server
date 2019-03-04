package groupware

import "testing"

func TestLogin(t *testing.T) {
	_, err := Login("hwkim03", "xogns00)(*")
	if err != nil {
		t.Fatal(err)
	}
}
