#!/bin/bash
mkdir -p ./var/log/golang && touch ./var/log/golang/info.log
docker-compose up -d 