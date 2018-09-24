package MusicLib

import "fmt"

type MusicEntry struct {
	ID       string
	Name     string
	Artist   string
	Source   string
	FileType string
}

func (m MusicEntry) ToString() string {
	return fmt.Sprintf("ID: %s ,Name: %s ,Artist: %s ,source: %s ,FileType: %s ,",
		m.ID, m.Name, m.Artist, m.Source, m.FileType)
}
