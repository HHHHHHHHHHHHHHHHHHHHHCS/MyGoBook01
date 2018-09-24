package MusicLib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed,not empty")
	}

	m0 := &MusicEntry{"1", "Will", "AAA", "www.baidu.com", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("NewMusicManager Add() is failed.")
	}

	m := mm.Find("Will")
	if m == nil {
		t.Error("NewMusicManager Find() is failed")
	}

	if m.ID != m0.ID || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.source != m0.source || m.FileType != m0.FileType {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}
}
