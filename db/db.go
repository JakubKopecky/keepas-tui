package db

import (
	"os"
	"strings"

	"github.com/tobischo/gokeepasslib"
)

type Entry struct {
	Title    string
	Username string
	Password string
	Url      string
	Notes    string
}

var Data []Entry

func OpenAndLoadDB(fileName string, pass string) {
	file, _ := os.Open(fileName)

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(pass)
	gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()

	for _, entry := range db.Content.Root.Groups[0].Entries {
		Data = append(Data, Entry{
			Title:    entry.GetTitle(),
			Password: entry.GetPassword(),
			Username: entry.GetContent("UserName"),
			Url:      entry.GetContent("URL"),
			Notes:    strings.ReplaceAll(entry.GetContent("Notes"), "\n", " "),
		})
	}
}
