#!/bin/zsh
echo "Running par-mnemonics"

GOGC=0   && go test -bench=. -benchtime=5x $PWD/par-mnemonics/ 
GOGC=50  && go test -bench=. -benchtime=5x $PWD/par-mnemonics/ 
GOGC=100 && go test -bench=. -benchtime=5x $PWD/par-mnemonics/ 
GOGC=200 && go test -bench=. -benchtime=5x $PWD/par-mnemonics/ 
