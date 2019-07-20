### Работа с ssh-gen



**Используемые утилиты**

- [Sops](https://github.com/mozilla/sops/releases)
- [Golang](https://golang.org/doc/install)

**Порядок работы**

1. Склонить репу
2. Выполнить команды  
    `sops -d credentials.json.enc > credentials.json`  
    `go run main.go < USERNAME > ~/.ssh/work.config`
3. Убедиться что в `~/.ssh/config` есть запись `Include ~/.ssh/work.config`

