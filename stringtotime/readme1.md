# Data string para time

No nosso dia a dia criamos progrmas go que podem receber alguma informação, imagine que neste caso se uma data no formato string, mas sua estrutura **(struct)** espresa receber um tipo time.Time ou seja data e hora. Como converter isso?

Aparentemente fazer essa converção é bem simples,mas não é tão simples assim, pois como é esperado um tipo "time" não da para so converter e pronto, temos que add um horário também.

## Exemplo

Vaja os exemplos abaixo

```go
type date struct {
	Date time.Time
}

func main() {
	var err error
	dte := date{} //criando a struct
	dateStr := "2019-08-31" // declarando variável com a data a ser convertida
	dte.Date, err = convertStringToTime(dateStr) //convertento...
	if err != nil {
		log.Println(fmt.Sprintf("date [%s], error: [%s]", dte.Date.Format("2006-01-02"), err))
	}
	log.Println("Date", dte.Date) // imprime no terminal a data convertida
}
```

Com esse bloco de código acima temos uma data do tipo *string* convertida para tipo *time*, para ver a função que usei na converção [clique aqui.](https://github.com/DiegoSantosWS/godate/blob/master/stringtotime/strtotime.go)


