package pkg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserDir(t *testing.T) {
	dirPath := MakeDir()
	filePath := fmt.Sprintf("%s/%s", dirPath, "t_t.txt")

	t.Run("making the dir", func(t *testing.T) {
		stat, err := os.Stat(dirPath)
		require.NoError(t, err)
		require.True(t, stat.IsDir())
	})

	t.Run("create file", func(t *testing.T) {
		stat, err := os.Stat(filePath)
		require.NoError(t, err)
		require.Equal(t, "t_t.txt", stat.Name())
	})

	t.Run("get user name", func(t *testing.T) {
		filename := "test.txt"
		username := "testy"
		// make file
		f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

		// write username
		f.Write([]byte(username))

		// call getUserName
		got := getUserName(filename)
		want := username
		// compare the two names
		require.Equal(t, got, want)
		// defer deleting all files
		defer f.Close()
		defer os.RemoveAll(filename)
	})
}
