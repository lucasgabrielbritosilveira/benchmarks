#!/bin/zsh
echo "Running par-mnemonics"

go test -bench=. -benchtime=10x $PWD/par-mnemonics/ 
