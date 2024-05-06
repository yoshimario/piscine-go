#!/bin/bash

# Fetch superhero data and filter by ID 70
name=$(curl -s https://01.gritlab.ax/assets/superhero/all.json | jq -r '.[] | select(.id == 70) | .name')

# Output the name
echo "\"name\"$name"