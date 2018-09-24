package MusicLib

import (
	"errors"
	"strconv"
)

type MusicManager struct {
	musics []MusicEntry
	id     int
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0), 1}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > m.Len() {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if m.Len() == 0 {
		return nil
	}

	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	music.ID = strconv.Itoa(m.id)
	m.musics = append(m.musics, *music)
	m.id++
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil
	}
	removedMusic := &m.musics[index]

	//从数组切片中删除元素
	if m.Len() == 1 { //删除只有一个的元素
		m.musics = make([]MusicEntry, 0)
	} else if index == 0 { //删除的是队列首领
		m.musics = m.musics[1:]
	} else if index == m.Len()-1 { //删除的是最后一个元素
		m.musics = m.musics[:index-1]
	} else if index < m.Len()-1 { //删除中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if name == "" {
		return nil
	}

	for k, v := range m.musics {
		if v.Name == name {
			return m.Remove(k)
		}
	}
	return nil
}
