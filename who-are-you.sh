#!/bin/bash

# Fetch superhero data and filter by ID 70
curl -s https://01.gritlab.ax/assets/superhero/all.json | jq '.[] | select ( .id == 70 )' | jq .name