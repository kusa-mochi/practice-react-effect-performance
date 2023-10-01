#!/bin/bash

curl --header "Content-Type: application/json" --data '{"name": "Kushiro"}' http://192.168.0.27:8080/greet.v1.GreetService/Greet
