#!/bin/bash

echo "$1" | docker run -i --rm --name script-python python python
