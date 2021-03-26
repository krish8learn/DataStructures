package src

import (
	"errors"
	"fmt"
)

type Song struct {
	SongName   string
	SongArtist string
	Previous   *Song
	Next       *Song
}

type PlayList struct {
	PlayListName string
	Head         *Song
	Tail         *Song
	NowPlaying   *Song
}

/*operation to perform: Create playlist, Add song, ShowAll songs, Start playing, Next/Previous song*/
//we will use Double Linked List

func CreatePlayList(name string) *PlayList {
	return &PlayList{
		PlayListName: name,
	}
	//other properties are by default nil
}

func CreateSong(name, artist string) *Song {
	return &Song{
		SongName:   name,
		SongArtist: artist,
	}
	//other properties are by default nil
}

func (p *PlayList) AddSong(name, artist string) error {
	NewSong := CreateSong(name, artist)

	if p.Head == nil {
		//this song is the first
		p.Head = NewSong
	} else {
		//this is not the 1st song
		currentNode := p.Tail
		currentNode.Next = NewSong
		NewSong.Previous = currentNode
	}
	p.Tail = NewSong
	return nil
}

func (p *PlayList) DeleteSong(name, artist string) error {
	temp := p.Head
	find := false
	if temp == nil {
		return errors.New("Empty Nothing to delete")
	}
	for temp.Next != nil {
		prev := temp.Previous
		ahead := temp.Next
		if temp.SongName == name && temp.SongArtist == artist {
			find = true
			prev.Next = temp.Next
			ahead.Previous = temp.Previous
			temp.Next = nil
			temp.Previous = nil
			break
		}
		temp = temp.Next
	}
	if find == false {
		return errors.New("No match found")
	}

	return nil
}

func (p *PlayList) ShowAll() error {
	temp := p.Head
	if temp == nil {
		return errors.New("Empty Playlist")
	}
	fmt.Printf("%v--->%v \n", *&temp.SongArtist, *&temp.SongName)
	for temp.Next != nil {
		temp = temp.Next
		fmt.Printf("%v--->%v \n", *&temp.SongArtist, *&temp.SongName)
	}
	return nil
}

func (p *PlayList) StartPlaying() *Song {
	p.NowPlaying = p.Head
	return p.NowPlaying
}

func (p *PlayList) NextSong() *Song {
	p.NowPlaying = p.NowPlaying.Next
	return p.NowPlaying
}

func (p *PlayList) PreviousSong() *Song {
	p.NowPlaying = p.NowPlaying.Previous
	return p.NowPlaying
}
