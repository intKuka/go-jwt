# JWT Authentication (partially)
Часть сервиса аутентификации, нацеленная на создание и обновление jwt-токенов.

## Запуск проекта
**Требования:**
 - MongoDB
 
**Шаги**
``` 
git clone git@github.com:intKuka/go-jwt.git
 ```
``` 
cd go-jwt
 ```
Нужно создать файл `.env`. В корневой папке есть пример: `.env.example`.

В `.env` нужно указать uri для подключения к MongoDB и придумать секретный ключ
``` 
go run main.go
 ```

## Работа с проетом
Порядок может быть такой:

1. Создать пользователя

   __POST /users__

2. Указать в параметре запроса существующий id пользователя. Результатом будет пара Access/Refresh Token

   __POST /auth/token__
   
### Refresh
По пути __POST /auth/refresh__ можно обновить пару токенов для конкретного пользователя.
В теле запроса в формате json нужно указать id пользователю и refresh токен, который был ему выдан
    
    {
        "refresh_token": "MjAyNC0wNC0wMyAwMzo1Mzo1OC44OTUzODIxICswMzAwIE1TSyBtPSszMC4zOTY2Njc5MDE=",
        "user_id": "446c1dc9-8c00-4adb-8e6d-52db5ac6f0a7"
    }

### Validate Access Token
Путь для проверки валидности токена __POST /auth/test__
