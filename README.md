### Работа с ssh-gen



**Используемые утилиты**

- [Sops](https://github.com/mozilla/sops/releases)
- [Golang](https://golang.org/doc/install)

**Порядок работы**

1. Склонить репу
2. Выполнить команды  
_При самом первом запуске `go run main.go`, не нужно перенаправлять в файл. Иначе не авторизируетесь на гугле._
   
    ```bash
    sops -d --output-type json credentials.json.enc
    go run main.go <USERNAME> > ~/.ssh/config.d/work
    ```
3. Убедиться что в `~/.ssh/config` есть запись `Include ~/.ssh/work.config`

