# Go Benchmarks
An adaptation of some Java benchmarks extracted from Dacapo and Reinassance suites to Go, with intuite to test Go Gargabe Collector.  

# Motivation


# Algorithms 

## par-mnemonics

A técnica de pair mnemonics utilizando streams do JDK para facilitar a memorização de pares de informações, como palavras ou números. Com
streams, você pode processar coleções de dados de forma declarativa e paralela. Cada
par é transformado em uma associação mnemônica significativa através de operações
com mapas

## xalan

Converte documentos XML, aplicando folhas de estilo XSLT para produzir
saídas formatadas. A escolha do Xalan se justifica por ser uma carga de trabalho customizável. Adapted from DaCapo suite

## fj-kmeans

O algoritmo de k-means utilizando Fork/Join em Java paraleliza o pro-
cesso de agrupamento de dados em clusters. Inicialmente, os pontos de dados são
distribuídos em subtarefas, cada uma responsável por calcular a distância dos pontos
aos centróides e atribuir os pontos ao cluster mais próximo. Cada subtarefa é tratada
como uma instância de uma classe que estende RecursiveTask. O ForkJoinPool
gerencia a execução paralela dessas subtarefas.