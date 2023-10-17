package main

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"omdbposter/omdbapi"
	"omdbposter/tui"
	"os"
	"reflect"
	"strings"
)

//var envFileFlag = flag.String("e", ".env", "dotenv file")

var buildFlag = flag.String("b", "", "build flags")

func init() {
	clearStdin()
}

func main() {
	flag.Parse()
	//err := godotenv.Load(*envFileFlag)
	//if err != nil {
	//	log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка загрузки файл %v", envFileFlag)))
	//}
	apiKey := os.Getenv("OMDBMOVIE_API_KEY")

	// Получение заголовка фильма
	var movieTitle string
	setTitleFromInput(&movieTitle)

	// Поиск фильмов, подходящих под введенное название
	var resultsOfSearch omdbapi.SearchResult
	search(&resultsOfSearch, &movieTitle, apiKey)

	// Выбор фильма из предложенных
	var selectedMovie omdbapi.MovieSearched
	for reflect.DeepEqual(selectedMovie, omdbapi.MovieSearched{}) {
		if isChange := setSelectedMovie(&selectedMovie, resultsOfSearch.Search); isChange {
			setTitleFromInput(&movieTitle)
			search(&resultsOfSearch, &movieTitle, apiKey)
		}
	}

	// TODO:
	movie, err := omdbapi.GetById(selectedMovie.ImdbID, apiKey)
	if err != nil {
		log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка получения фильма: %v", err)))
	}

	pagerProgram := tea.NewProgram(
		tui.PagerModel{
			Movie: movie,
		},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse )
	)
	if _, err := pagerProgram.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

func search(results *omdbapi.SearchResult, movieTitle *string, apiKey string) {
	var err error
	*results, err = omdbapi.SearchByTitle(*movieTitle, apiKey)
	if err != nil {
		log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка поиска: %v", err)))
	}
	for results.TotalResults == 0 {
		clearStdin()
		fmt.Print("\nПо Вашему запросу фильмов не найдено 😞\nПопробуйте еще раз\n\n")
		setTitleFromInput(movieTitle)
		*results, err = omdbapi.SearchByTitle(*movieTitle, apiKey)
		if err != nil {
			log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка поиска: %v", err)))
		}
	}
}

func setTitleFromInput(movieTitle *string) {
	inputProgram := tea.NewProgram(tui.RunInput(
		"Forrest Gump...", "Введите название фильма: "))
	m, err := inputProgram.Run()
	if err != nil {
		log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка ввода фильма: %v", err)))
	}
	if m.(tui.InputModel).IsExit {
		os.Exit(0)
	}
	*movieTitle = m.(tui.InputModel).TextInput.Value()
}

// return true если нужно изменить название
func setSelectedMovie(selected *omdbapi.MovieSearched, movies []omdbapi.MovieSearched) bool {
	selectProgram := tea.NewProgram(tui.RunSelect(movies, "Найдено по Вашему запросу: "))
	selectModel, err := selectProgram.Run()
	if err != nil {
		log.Fatal(strings.ToUpper(fmt.Sprintf("ошибка выбора: %v", err)))
	}

	if m, ok := selectModel.(tui.SelectModel); ok {
		switch m.SelectCode {
		case -1:
			os.Exit(0)
		case 0:
			clearStdin()
			return true
		case 1:
			*selected = m.Choice
		}
	}
	return false
}

func clearStdin() {
	fmt.Print("\033[H\033[2J")
}
