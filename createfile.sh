#!/bin/bash

#create the two required files
if [ ! -f "reqfile.log" ]; 
then
touch reqfile.log
fi

if [ ! -f "error.log" ];
then 
touch error.log
fi