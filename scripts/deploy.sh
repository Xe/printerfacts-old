#!/bin/sh

heroku container:push -a printerfacts web
heroku container:release -a printerfacts web
