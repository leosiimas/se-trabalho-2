---
sticker: emoji//1f4ca
---
# Trabalho 2

Este trabalho tem como objetivo criar três aplicações que realizam a mesma tarefa, mas com maneira diferentes. Para cada aplicação será feito em três linguagens diferentes, escolhidas: *C*, *Python* e *Go*. 

Os algoritmos escolhidos são: *Bucket Sort*, *Counting Sort* e *Radix Sort*. Esses são três algoritmos de ordenação que executam a mesma tarefa de maneiras diferentes.

## Execução

Para executar os arquivos precisa apenas executar o *Makefile* na raiz do projeto
```shell
comando
```

Os algoritmos podem ser executados separadamente com os seguintes comando
**![![UFRGS-2023-1/SE/trabalho-2/#^Table]]



## Algoritmos

Antes de mostrar os códigos e os resultados dos testes, é importante entender como os algoritmos funcionam, e como cada um aborda o mesmo problema de forma diferente, quais são suas vantagens e desvantagens entre si.

### *Bucket Sort*
 *Bucket Sort* ou *bin sort* é um algoritmo de ordenação que funciona dividindo um vetor em um número finito de recipientes. Cada recipientes é então ordenado individualmente, usando um algoritmo de ordenação diferente, ou usando o algoritmo *bucket sort* recursivamente. 

*Complexidade e Desempenho*:
- Pior caso: O(n<sup>2</sup>) => Ocorre quando todos os elementos são armazenados no mesmo balde considerando que o algoritmo usado para classificação seria por *ordenação por inserção*
- Médio: O(n + k) => 
- Melhor caso: O(n+ k) =>

*Funcionamento*
1. Inicia um vetor de "baldes", inicialmente vazios
2. Va para o vetor original, incluindo cada elemento em um balde
3. Ordene todos os baldes não vazios
4. Coloque os elementos dos baldes que não estão vazios no vetor original

![[Pasted image 20230807233454.png]]

### *Counting Sort*

### *Radix Sort*

## Linguagens

### *C*

### *Python*

### *Go*

#### Instalando Go

Para utilizar *Golang* no *Ubuntu*, bas apenas executar no terminal:
```shell
sudo apt-get update && sudo apt-get -y install golang-go
```

Checando a versão:
```shell
go version
```

E para executar qualquer programa em *Golang*:
```shell
go run [arquivo].go
```



## 


## Referencias

### *Bucket Sort*
- https://pt.wikipedia.org/wiki/Bucket_sort
- https://en.wikipedia.org/wiki/Bucket_sort
- https://gostudent.github.io/Letsgo/implementation-of-bucket-sort-in-GO
- https://www.programiz.com/dsa/bucket-sort

### *Counting Sort*
- https://www.programiz.com/dsa/counting-sort
- https://pt.wikipedia.org/wiki/Counting_sort
- https://en.wikipedia.org/wiki/Counting_sort

### *Radix Sort*
- https://www.programiz.com/dsa/radix-sort
- https://pt.wikipedia.org/wiki/Radix_sort
- https://en.wikipedia.org/wiki/Radix_sort
- https://www.geeksforgeeks.org/radix-sort/

