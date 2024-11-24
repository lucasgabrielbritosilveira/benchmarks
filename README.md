# Go Benchmarks
An adaptation of some Java benchmarks extracted from Dacapo and Reinassance suites to Go, with intuite to test Go Gargabe Collector.  

# Motivation


# Algorithms 

## par-mnemonics

The pair mnemonics technique uses JDK streams to facilitate the memorization of pairs of information, such as words or numbers. With
streams, you can process collections of data declaratively and in parallel. Each
pair is transformed into a meaningful mnemonic association through operations
with maps

## xalan

Converts XML documents by applying XSLT stylesheets to produce
formatted outputs. The choice of Xalan is justified because it is a customizable workload. Adapted from DaCapo suite

## fj-kmeans

The k-means algorithm using Fork/Join in Java parallelizes the pro-
process of grouping data into clusters. Initially, the data points are
distributed in subtasks, each responsible for calculating the distance of the points
to the centroids and assigning the points to the closest cluster. Each subtask is handled
as an instance of a class that extends RecursiveTask. The ForkJoinPool
manages the parallel execution of these subtasks.