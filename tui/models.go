package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"omdbposter/omdbapi"
)

type SelectModel struct {
	Header  string
	cursor  int
	Choice  omdbapi.MovieSearched
	choices []omdbapi.MovieSearched
	// 1 - пользователь выбрал
	// 0 - изменение название фильма
	// -1 - выход их программы
	SelectCode int
}

type PagerModel struct {
	Movie    omdbapi.Movie
	ready    bool
	viewport viewport.Model
}

type InputModel struct {
	Header    string
	TextInput textinput.Model
	IsExit    bool
	err       error
}
