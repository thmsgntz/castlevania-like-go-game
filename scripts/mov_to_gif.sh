#!/bin/bash

# https://superuser.com/questions/436056/how-can-i-get-ffmpeg-to-convert-a-mov-to-a-gif
# https://stackoverflow.com/questions/34341808/is-there-a-way-to-add-a-gif-to-a-markdown-file

set +x

MOVIES_DIRECTORY="${HOME}/Documents/screen_recordings"

file_mov=$(ls -t ${MOVIES_DIRECTORY}/*.mov | head -n1)
echo $file_mov

ffmpeg -ss 00:00:00.000 -i "${file_mov}" -pix_fmt rgb24 -r 10 -s 800x600 -t 00:00:10.000 ressources/gifs/output.gif 

convert -layers Optimize ressources/gifs/output.gif ressources/gifs/game_animations.gif && rm -f ressources/gifs/output.gif

du -h ressources/gifs/game_animations.gif