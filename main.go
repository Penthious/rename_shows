package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"regexp"
)

const (
	api  string = "https://api.tvmaze.com"
	path string = "D:/Videos/Rename Shows"
)

func main() {
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
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

	fmt.Println("Type the name of your show: ")
	fmt.Println("Example: Rick and Morty (this needs to be the same name as the folder your working on)")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading string: %v\n", err)
		return
	}
	re := regexp.MustCompile(`\r?\n`)
	name = re.ReplaceAllString(name, "")

	fmt.Println("Type in your path to the shows: ")
	fmt.Println("Example: C:\\Users\\Beast\\projects\\rename_shows")
	pathPtr, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading string: %v\n", err)
		return
	}
	pathPtr = re.ReplaceAllString(pathPtr, "")

	showDirectory := fmt.Sprintf("%s/%s", pathPtr, name)
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

	fmt.Println("\n\n\n\nAll done: click X or type `ctrl+c`")
}
