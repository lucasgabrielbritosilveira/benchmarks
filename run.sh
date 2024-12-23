#!/bin/zsh

parmnemonics=true
xalan=true
fj_kmeans=false

if [ $parmnemonics = true ]; then

echo "Running par-mnemonics"

echo "GOGC,Mallocs,Frees,HeapAlloc,HeapSys,HeapReleased,NextGC,LastGC,PauseTotalNs,NumGC,NumForcedGC,GCCPUFraction" > ./output/parmnemonics_results.csv  

cd benchmarks/

export GOGC=25  && go test -bench=. -benchtime=10x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=50  && go test -bench=. -benchtime=10x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=100 && go test -bench=. -benchtime=10x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=200 && go test -bench=. -benchtime=10x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 
export GOGC=400 && go test -bench=. -benchtime=10x $PWD/par-mnemonics/  && cat $PWD/par-mnemonics/results.csv >> ../output/parmnemonics_results.csv 

rm $PWD/par-mnemonics/results.csv
cd ../

fi 

if [ $xalan = true ]; then

echo "Running xalan"

echo "GOGC,Mallocs,Frees,HeapAlloc,HeapSys,HeapReleased,NextGC,LastGC,PauseTotalNs,NumGC,NumForcedGC,GCCPUFraction" > ./output/xalan_results.csv  

cd benchmarks/

export GOGC=25  && go test -bench=. -benchtime=10x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=50  && go test -bench=. -benchtime=10x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=100 && go test -bench=. -benchtime=10x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=200 && go test -bench=. -benchtime=10x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 
export GOGC=400 && go test -bench=. -benchtime=10x $PWD/xalan/  && cat $PWD/xalan/results.csv >> ../output/xalan_results.csv 

rm $PWD/xalan/results.csv

cd ../
fi