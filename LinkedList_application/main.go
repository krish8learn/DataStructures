package main

import (
	"fmt"

	"github.com/krish8learn/DataStructures/LinkedList_application/src"
)

func main() {
	playlistname := "MyPlayList"

	MyPlayList := src.CreatePlayList(playlistname)
	fmt.Println("PlayList Created")
	fmt.Println("ADDING SONGS TO THE PLAYLIST")
	MyPlayList.AddSong("Numb", "Linkin Park")
	MyPlayList.AddSong("Locking up the Sun ", "Poets of the Fall")
	MyPlayList.AddSong("ChopSuey", "System of down")
	MyPlayList.AddSong("5th Symphony", "Bethoven")
	MyPlayList.AddSong("Breaking The Habit", "Linkin Park")
	MyPlayList.AddSong("Cradle for love", "Poets of the fall")
	err := MyPlayList.ShowAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	/*MyPlayList.StartPlaying()
	fmt.Printf("Now : %s by %s \n", MyPlayList.NowPlaying.SongName, MyPlayList.NowPlaying.SongArtist)

	MyPlayList.NextSong()
	fmt.Printf("NextSong : %s by %s\n", MyPlayList.NowPlaying.SongName, MyPlayList.NowPlaying.SongArtist)

	MyPlayList.PreviousSong()
	fmt.Printf("Previous Song : %s by %s\n", MyPlayList.NowPlaying.SongName, MyPlayList.NowPlaying.SongArtist)*/

	MyPlayList.DeleteSong("ChopSuey", "System of down")
	Merr := MyPlayList.ShowAll()
	if Merr != nil {
		fmt.Println(Merr)
	}
	fmt.Println()
}
