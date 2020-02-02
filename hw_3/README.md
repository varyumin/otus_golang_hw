# TO DO
## Распаковка строки
Домашнее задание
Частотный анализ
Цель: Напиcать функцию, принимающую на вход строку с текстом и возвращающую слайс с 10 самыми частовстречающимеся в тексте словами.
Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать функцию вида Top10(string) ([]string)
При необходимости выделить вспомогательные функции
Написать unit-тесты на функцию

Пояснения

Если есть более 10 самых частотых слов (например 15 разных слов встречаются ровно 133 раза, остальные < 100), можно вернуть любые 10 из самых частотных.

Словоформы не учитываем. "нога", "ногу", "ноги" - это разные слова.
Слово с большой и маленькой буквы можно считать за разные слова. "Нога" и "нога" - это разные слова.
Знаки препиания можно считать "буквами" слова или отдельными словами. "-" (тире) - это отдельное слово. "нога," и "нога" - это разные слова.

Пример: "cat and dog one dog two cats and one man". "dog", "one", "and" - встречаются два раза, это топ-3.

Задание со звездочкой (*): учитывать большие/маленьгие буквы и знаки препинания. "Нога" и "нога" - это одинаковые слова, "нога," и "нога" - это одинаковые слова, "—" (тире) - это не слово.
Критерии оценки: Функция должна проходить все тесты
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и проверить пакет с помощью go get / go test

# HOW TO
## Arguments:
```
Usage of of hw_3:
  -file string
        Path to text file
  -line string
        String with values
  -top int
        Number of top (default 10)

```

## RUN
```
git clone git@github.com:varyumin/otus_golang_hw.git
cd otus_golang_hw/hw_3/
go mod vendor
go test

go run main.go -line={WRAPPER_LINE}
OR
go run main.go -file={PATH_TO_FILE}

```

### OR

```
make run
```

## Example:
```
$FILE=voyna-i-mir-tom-1.txt make run

go: finding golang.org/x/lint latest
go: finding golang.org/x/lint/golint latest
go: finding golang.org/x/tools latest
=== RUN   TestUnpack
--- PASS: TestUnpack (0.00s)
    main_test.go:26: "aaaabccddddde" == "aaaabccddddde"
    main_test.go:26: "abcd" == "abcd"
    main_test.go:26: "" == ""
PASS
ok      github.com/varyumin/otus_golang_hw/hw_2 (cached)

gggg
```
### OR

```
$go run main.go -line="a s f g k j y h g f g h n g f d s a s d f v b g h j u y t r e w e e w q q w e r t t"
 1 place is taken by the word: 'g' occurs 5 times
 2 place is taken by the word: 'e' occurs 4 times
 3 place is taken by the word: 'f' occurs 4 times
 4 place is taken by the word: 't' occurs 3 times
 5 place is taken by the word: 'w' occurs 3 times
 6 place is taken by the word: 's' occurs 3 times
 7 place is taken by the word: 'h' occurs 3 times
 8 place is taken by the word: 'j' occurs 2 times
 9 place is taken by the word: 'q' occurs 2 times
 10 place is taken by the word: 'a' occurs 2 times
```