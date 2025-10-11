#!/bin/bash

#create the two required files
if [ ! -f "req.log" ]; 
then
touch req.log
fi

if [ ! -f "error.log" ];
then 
touch error.log
fi