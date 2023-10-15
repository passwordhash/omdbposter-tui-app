# Omdbposter 

Консольное приложение на GO для быстрого поиска фильма и получения небольшой справки о нем.

![logo](demonstration.gif) 

## Установка

### MacOS (build из исходников)

- `git clone ...`
- `export OMDBMOVIE_API_KEY={{YOUR_API_KEY}}` Можете воспользоваться моим: *314bf63c*
- `go build -ldflags "-X main.omdbmovie_api_key=$OMDBMOVIE_API_KEY"`

Запуск: `./omdbposter`

### MacOS (builded)

1. Скачать файл **omdbposter**
2. Запустить: `./omdbposter`
