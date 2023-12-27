package pkg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserDir(t *testing.T) {
	path := dirPath()
	testfilePath := fmt.Sprintf("%s/%s", path, "t_t.txt")
	username := "testy"

	os.Mkdir(path, 0750)
	os.OpenFile(testfilePath, os.O_RDWR|os.O_CREATE, 0755)

	// tests getUserName function
	t.Run("get user name", func(t *testing.T) {
		// make file
		f, _ := os.OpenFile(testfilePath, os.O_RDWR|os.O_CREATE, 0755)

		// write username
		f.Write([]byte(username))

		// call getUserName
		got := getUserName(testfilePath)
		want := username
		// compare the two names
		require.Equal(t, got, want)
		// defer deleting all files
		defer f.Close()
		defer os.RemoveAll(path)

	})
}
