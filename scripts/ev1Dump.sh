#!/bin/bash
HOME=/home/jmainguy/Nextcloud/proxmark
if [ "$1" = "" ]
then
  echo "Usage: $0 first argument is card name"
  exit 1
fi

FILENAME=$1

pm3 -c "hf mfu dump -f ${FILENAME}"
