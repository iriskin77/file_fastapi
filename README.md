## 1) Что это такое?

Приложение для определения тональности текста. Приложение позволяет загружать, скачивать и
обрабатывать с помощью ML модели файлы в формате excel. 

## 2) Основные технологии

+ FastApi
+ PostgreSQL
+ Tortoise ORM
+ Keras, Pymystem
+ Pandas

Подробнее см. файл requirements.txt

## 3) Как это работает?

Пользователь загружает файл в формате excel с текстами, которые необходимо обработать.

+ Предварительно файл необходимо подготовить. Шапка таблицы с названиями колонок должна распологаться в первой строчке таблицы, как на рисунке 1 (см. рис.1).
    Если этого не сделать, то программа не сможет обработать комментарии


![Рис. 1](/images/file_excel.png)

+ После этого необходимо загрузить файл и точно указать название колонки, в которой
находятся тексты. Предварительно нужно зарегистрироваться и авторизоваться, см. рис. 2



![Рис. 2](/images/upload.png)

После этого можно обработать файл (/process_file), указав его id

## 4) Как установить?

Приложение развернуто на хостинге: https://textemotions.onrender.com/docs



