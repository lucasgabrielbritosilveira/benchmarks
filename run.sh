#!/bin/zsh

parmnemonics=true
xalan=true
fjkmeans=true

header="GOGC,Mallocs,Frees,HeapAlloc,HeapSys,HeapReleased,NextGC,LastGC,PauseTotalNs,NumGC,NumForcedGC,GCCPUFraction,TotalTimeExecutionInNanoseconds"

if [ $parmnemonics = true ]; then

echo "Running par-mnemonics"

echo $header > ./output/parmnemonics_results.csv  
cd benchmarks/

export GOGC=50   && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=100  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=150  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=200  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=250  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=300  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=350  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=400  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=450  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=500  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=550  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=600  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=650  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=700  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=750  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=800  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=850  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=900  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=950  && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=1000 && go test -bench=. -benchtime=100x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 


rm $PWD/par-mnemonics/results.csv
cd ../

fi 

if [ $xalan = true ]; then

echo "Running xalan"

echo $header > ./output/xalan_results.csv  

cd benchmarks/

export GOGC=50   && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=100  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=150  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=200  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=250  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=300  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=350  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=400  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=450  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=500  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=550  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=600  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=650  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=700  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=750  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=800  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=850  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=900  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=950  && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=1000 && go test -bench=. -benchtime=100x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 


rm $PWD/xalan/results.csv

cd ../
fi

if [ $fjkmeans = true ]; then

echo "Running fj-kmeans"

echo $header > ./output/fjkmeans_results.csv  

cd benchmarks/

export GOGC=50   && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=100  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=150  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=200  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=250  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=300  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=350  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=400  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=450  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=500  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=550  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=600  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=650  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=700  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=750  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=800  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=850  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=900  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=950  && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 
export GOGC=1000 && go test -bench=. -benchtime=100x $PWD/fj-kmeans/  && cat $PWD/fj-kmeans/results.csv >> ../output/fjkmeans_results.csv 


rm $PWD/fj-kmeans/results.csv

cd ../
fi