#!/bin/bash
docker build . -t code-quality || exit
docker run -it --volume /Users/ganesh/go/src/github.com/carousell/Feedback:/app  code-quality:latest findlanguage