#!/bin/bash
#find . | wc -l!/bin/sh
find . | wc -l | sed 's/ //g'