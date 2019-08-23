# Transformando data e hora atual em string

### Neste exemplo vamos transformar data em string

Em Go isso é muito simples, pois já foi pensado para ser simples.
Estamos acostumado a usar caractres para definir como será o posição da dos itens, em PHP se faz da seguinte forma

```php
$today = date("Y-m-d");
echo "Date ".$today;
/***
~$ 2019-08-22
***/
```

Em go isso é bem diferente, para pegar a data atua, usamos o pacote **time** da propría linguagem.

Exemplo:

```go

import (
    "time"
    "fmt"
)

func main() {
    now := time.Now()
}
```

A variavel **now** em sua tradução literal é o agora, sendo assim temos o time "tempo" corrente, com essas informação já podemos seguir. O codigo assim se executado vai gerar um erro, por que não podemos declarar uma variavel e não usa lá.

Seguinte nã func main... vamos ver como pegamos a data atual e exibimos.

```go
dateStr := now.Format("2006-01-02")
fmt.Println("Data atual é:", dateStr)

// ~$ Data atual é: 2019-08-22 *exibição*
```

Ah mas por que você usou a variaval **now**?
Bom como time é um pacote ele tem uma series de funções publícas e uma delas é **Format**, essa função recebe um parametro que é uma data na minha opnião é "locura" mas já acostumei e decorei kkkkkk, e com esse
parametro definimos o formato que a data será exposta na tela ou terminal. Como é feito no exemplo acima.

Ainda usando a função **Format** vamos mostrar a data e hora atual, que também é bem simples, basta add a hora **15:04:02** logo após a data.

```go
dateHoraStr := now.Format("2006-01-02 15:02:04")
fmt.Println("Data atual é:", dateHoraStr)

// ~$ Data e hora atual é: 2019-08-22 22:08:20 *exibição*
```

Por hoje é isso espero ter ajudado.

Para ver o exemplo formatado eu acesse o arquivo [string.go](https://github.com/DiegoSantosWS/godate/blob/master/datetostring/string.go)
