## Переделать по-новому:

Новая схема:

![Рис. 2](/images/new_scema.png)

1) Api-gateway на Go

2) Auth-service на Go + Mongo

3) File-service на Python, gRPC

4) File-process-service на Python

## Оглавление

+ [1) Описание?](#title1)
+ [2) Как это работает](#title2)
+ [3) Как установить](#title3)
+ [4) Заметки](#title4)

## <a id="title1">1) Описание</a>

Приложение для определения тональности текста. Приложение позволяет загружать, скачивать и
обрабатывать с помощью ML модели файлы в формате excel. 


Схема БД:

![Рис. 1](/images/db_schema.png)

## <a id="title2">2) Основные технологии</a>

+ FastApi
+ PostgreSQL
+ Tortoise ORM
+ Keras, Pymystem
+ Pandas
+ Docker, docker-compose

Подробнее см. файл requirements.txt

## <a id="title3">3) Как это работает?</a>

Пользователь загружает файл в формате excel с текстами, которые необходимо обработать.

1) Предварительно файл необходимо подготовить. Шапка таблицы с названиями колонок должна распологаться в первой строчке таблицы, как на рисунке 1 (см. рис.1).
    Если этого не сделать, то программа не сможет обработать комментарии (см. рис. 2).


![Рис. 2](/images/file_excel.png)

2) Необходимо зарегистрироваться и авторизоваться (авторизация происходит по JWT), см. рис. 3

![Рис. 3](/images/reg_user.png)

3) После этого необходимо загрузить файл и точно указать название колонки, в которой
находятся тексты (см. рис. 4).

   
![Рис. 4](/images/upload.png)

После этого можно обработать файл (/process_file), указав его id (см. рис. 5)

![Рис. 5](/images/file_process.png)

Пример обработанных комментариев (см. рис 6):

![Рис. 6](/images/comments.png)

## <a id="title4">4) Как установить?</a>

Клонировать репозиторий: https://github.com/iriskin77/api_lstm_model.git

Запустить приложение из Docker:
 + docker-compose build
 + docker-compose up
