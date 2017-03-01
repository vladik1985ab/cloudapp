package model

type BuildpackCustom struct {
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	Locked   bool   `json:"locked"`
	Filename string `json:"filename"`
}

