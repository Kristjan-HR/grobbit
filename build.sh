#! /usr/bin/env bash

rm ./a.out ./Grobbit ./out/test_student.ll ./out/test_student.ll_norab
go build Grobbit
./Grobbit
clang -O3 out/main.c out/test_student.ll
