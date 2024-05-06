#!/bin/bash

curl -s  https://01.gritlab.ax/assets/superhero/all.json | jq --arg HERO_ID "$HERO_ID" -r '.[] | select(.id == ($HERO_ID|tonumber)) | .connections.relatives' | tr -d '"'