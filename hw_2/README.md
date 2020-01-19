# TO DO
## Распаковка строки
Цель: Создать Go функцию, осуществляющую примитивную распаковку строки,  
содержащую повторяющиеся символы / руны, например: 
* "a4bc2d5e" => "aaaabccddddde" 
* "abcd" => "abcd" 
* "45" => "" (некорректная строка) 
* "" => ""  
Дополнительное задание: поддержка escape - последовательностей 
* `qwe\4\5` => `qwe45` (*) 
* `qwe\45` => `qwe44444` (*) 
* `qwe\\5` => `qwe\\\\\` (*)  
В случае если была передана некорректная строка функция должна возвращать ошибку.  
Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать функцию вид Unpack(string) (string, error)
При необходимости выделить вспомогательные функции
Написать unit-тесты на функцию  

Критерии оценки: 
- Функция должна проходить все тесты
- Код должен проходить проверки go vet и golint
- У преподавателя должна быть возможность скачать и проверить пакет с помощью go get / go test

# HOW TO
## Arguments:
```
Usage of hw_2:
  -line string
    	primitive unboxing of a string containing duplicate characters/runes

```

## RUN
```
git clone git@github.com:varyumin/otus_golang_hw.git
cd otus_golang_hw/hw_2/
go mod vendor
go test
go run main.go -line={WRAPPER_LINE}
```

### OR

```
make run
```

## Example:
```
$LINE=g4 make run

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
$go run main.go -line=g4
gggg
```