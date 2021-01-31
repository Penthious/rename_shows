package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/manifoldco/promptui"
)

type ShowsInterface interface {
	GetPossibleShows(showName string, api string) (ShowsStruct, error)
}

type ShowInterface interface {
	PickShow(shows ShowsStruct) (ShowStruct, error)
}

func NewShow() ShowInterface {
	return ShowStruct{}
}
func NewShows() ShowsInterface {
	return ShowsStruct{}
}

func (s ShowsStruct) GetPossibleShows(showName string, api string) (ShowsStruct, error) {
	resp, err := http.Get(fmt.Sprintf("%s/search/shows?q=%s", api, showName))
	if err != nil {
		return ShowsStruct{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ShowsStruct{}, err
	}

	var shows ShowsStruct
	err = json.Unmarshal(body, &shows)
	if err != nil {
		return ShowsStruct{}, err
	}

	return shows, nil
}

func (s ShowStruct) PickShow(shows ShowsStruct) (ShowStruct, error) {
	var showNames []string
	for _, show := range shows {
		showNames = append(showNames, show.ShowStruct.Name)
	}

	prompt := promptui.Select{
		Label: "Select which show/movie that matches",
		Items: showNames,
	}

	index, _, err := prompt.Run()

	if err != nil {
		return ShowStruct{}, err
	}

	return shows[index].ShowStruct, nil
}
