#!/bin/bash
HOME=/home/jmainguy/Nextcloud/proxmark
if [ "$1" = "" ]
then
  echo "Usage: $0 first argument is card name"
  echo "Example: $0 ev17"
  exit 1
fi

FILENAME=${1}

## Dump Card
ev1Dump.sh $FILENAME

# Display a prompt to the user
read -p $'\e[1;32mReplace card on reader with clone and press Enter:\e[0m '

# Continue with the rest of the script
echo "Continuing with the script..."

ev1Restore.sh /home/jmainguy/Nextcloud/proxmark/${FILENAME}.json
