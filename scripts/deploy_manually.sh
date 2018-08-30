#!/bin/sh

docker build -t registry.heroku.com/printerfacts/web .
docker push registry.heroku.com/printerfacts/web
echo "run heroku container:release -a printerfacts web"
