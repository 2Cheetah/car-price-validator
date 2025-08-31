package repository

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Filters struct {
	Data Data `json:"data"`
}

type Data struct {
	Type       string     `json:"type"`
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Make  MakeRoot  `json:"make"`
	Model ModelRoot `json:"model"`
}

type MakeRoot struct {
	Values []Make `json:"values"`
}

type Make struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ModelRoot struct {
	Values []ModelParent `json:"values"`
}

type ModelParent struct {
	ParentID string     `json:"parent_id"`
	Values   []CarModel `json:"values"`
}

type CarModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func ReadData(path string) (Data, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return Data{}, fmt.Errorf("couldn't read file with filters, error: %w", err)
	}

	var filters Filters
	if err := json.Unmarshal(f, &filters); err != nil {
		return Data{}, fmt.Errorf("couldn't unmarshal file with filters, error: %w", err)
	}
	slog.Debug("filters.json", "type", filters.Data.Type, "id", filters.Data.Id)

	return filters.Data, nil
}

func GetCarMakes() ([]string, error) {
	data, err := ReadData("./.temp/filters.json")
	if err != nil {
		return nil, fmt.Errorf("couldn't get data, error: %w", err)
	}

	carMakes := make([]string, 0)
	for _, parent := range data.Attributes.Make.Values {
		carMake := strings.ToLower(parent.Name)
		carMakes = append(carMakes, carMake)
	}
	return carMakes, nil
}

func GetCarMakesData(data Data) map[string]string {
	makeData := make(map[string]string, 0)
	for _, parent := range data.Attributes.Make.Values {
		carMake := strings.ToLower(parent.Name)
		makeData[carMake] = parent.ID
	}
	return makeData
}

func GetCarModelsByMake(makeName string) ([]string, error) {
	data, err := ReadData("./.temp/filters.json")
	if err != nil {
		return nil, fmt.Errorf("couldn't get data, error: %w", err)
	}

	makeData := GetCarMakesData(data)
	makeName = strings.ToLower(makeName)
	makeID, ok := makeData[makeName]
	if !ok {
		return nil, fmt.Errorf("make %q not found", makeName)
	}

	for _, parent := range data.Attributes.Model.Values {
		if parent.ParentID != makeID {
			continue
		}
		carModels := make([]string, 0, len(parent.Values))
		for _, model := range parent.Values {
			carModel := strings.ToLower(model.Name)
			carModels = append(carModels, carModel)
		}
		return carModels, nil
	}

	return []string{}, fmt.Errorf("no car models found for that make")
}
