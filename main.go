package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	api  string = "https://api.tvmaze.com"
	path string = "D:/Videos/Rename Shows"
)

func main() {
	go forever()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

}

func forever() {
	var name string
	e := NewEpisodes()
	s := NewShow()
	ss := NewShows()

	fmt.Print("Type the name of your show: ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading string: %v\n", err)
		return
	}
	name = strings.ReplaceAll(name, "\r", "")
	name = strings.ReplaceAll(name, "\n", "")

	pathPtr := flag.String("path", path, "The name of the show that you want to rename")
	flag.Parse()

	showDirectory := fmt.Sprintf("%s/%s", *pathPtr, name)
	folders, err := ioutil.ReadDir(showDirectory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	shows, err := ss.GetPossibleShows(name, api)
	if err != nil {
		fmt.Printf("Error getting shows: %v\n", err)
		return
	}

	show, err := s.PickShow(shows)
	if err != nil {
		fmt.Printf("Error picking show: %v\n", err)
		return
	}

	episodes, err := e.GetEpisodes(show, api)
	if err != nil {
		fmt.Printf("Error getting episods: %v\n", err)
		return
	}

	err = e.UpdateFiles(folders, episodes, showDirectory, name)
	if err != nil {
		fmt.Printf("Error updating files: %v", err)
		return
	}

}
