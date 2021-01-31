package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Episodes []struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Number  int    `json:"number"`
	Summary string `json:"summary"`
}

func NewEpisodes() EpisodesInterface {
	return Episodes{}
}

type EpisodesInterface interface {
	GetEpisodes(show ShowStruct, api string) (Episodes, error)
	UpdateFiles(folders []os.FileInfo, episodes Episodes, path string, showName string) error
	convertEpisodes(episodes Episodes, files []os.FileInfo, path string, showName string) error
	filterEpisodes(episodes Episodes, season string) Episodes
}

func (e Episodes) GetEpisodes(show ShowStruct, api string) (Episodes, error) {
	resp, err := http.Get(fmt.Sprintf("%s/shows/%v/episodes", api, show.ID))
	if err != nil {
		return Episodes{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Episodes{}, err
	}

	var episodes Episodes
	err = json.Unmarshal(body, &episodes)
	if err != nil {
		return Episodes{}, err
	}

	return episodes, nil

}

func (e Episodes) UpdateFiles(folders []os.FileInfo, episodes Episodes, path string, showName string) error {

	for _, f := range folders {
		season := strings.Split(f.Name(), " ")[len(folders)-1]
		filteredEpisodes := e.filterEpisodes(episodes, season)

		files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", path, f.Name()))
		if err != nil {
			return err
		}

		if err := e.convertEpisodes(filteredEpisodes, files, fmt.Sprintf("%v/%v", path, f.Name()), showName); err != nil {
			return err
		}

	}
	return nil
}

func (e Episodes) convertEpisodes(episodes Episodes, files []os.FileInfo, path string, showName string) error {
	season := FixNumbers(episodes[0].Season)

	if len(episodes) >= len(files) {
		fmt.Print("Episodes dont match\n")
		var accept string
		fmt.Print("Do you want to Continue? Yes/No ")
		fmt.Scanf("%s", &accept)
		if strings.Contains(strings.ToLower(accept), "no") {
			return errors.New("Exited by user")
		}

	}

	fmt.Printf("\n================================================\n")
	fmt.Printf("================================================\n")
	fmt.Printf("================================================\n")
	fmt.Printf("Season being worked on: %v\n", season)
	for i, episode := range episodes {
		if len(files) <= i {
			break
		}
		episodeNumber := FixNumbers(episode.Number)
		fileArray := strings.Split(files[i].Name(), ".")
		ext := fileArray[len(fileArray)-1]
		fileName := fmt.Sprintf("%v - s%ve%v - %v.%v", showName, season, episodeNumber, episode.Name, ext)
		originalPath := fmt.Sprintf("%v/%v", path, files[i].Name())
		newPath := fmt.Sprintf("%v/%v", path, fileName)

		fmt.Printf("\nChanging %v\n", originalPath)
		fmt.Printf("To       %v\n", newPath)

		if err := os.Rename(originalPath, newPath); err != nil {
			return err
		}

	}

	return nil
}

func (e Episodes) filterEpisodes(episodes Episodes, season string) Episodes {

	var filtered Episodes

	for _, episode := range episodes {

		s := strconv.Itoa(episode.Season)
		if strings.Contains(season, s) {
			filtered = append(filtered, episode)
			// *episodes = append(episodes[:index], episodes[index+1]...)
		}
	}

	return filtered
}
