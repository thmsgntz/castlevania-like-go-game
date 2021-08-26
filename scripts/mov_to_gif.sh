#!/bin/bash

# https://superuser.com/questions/436056/how-can-i-get-ffmpeg-to-convert-a-mov-to-a-gif
# https://stackoverflow.com/questions/34341808/is-there-a-way-to-add-a-gif-to-a-markdown-file

set -x

MOVIES_DIRECTORY="${HOME}/Documents/screen_recordings"
OUTPUT_DIRECTORY="ressources/gifs/"
OUTPUT_GIF_NAME="game_animations.gif"
FFMEPG_LOGLEVEL=0
GIF_SIZE=800x600

d=$(date +"%d-%m-%Y-%H-%M-%S")
mv -f ${OUTPUT_DIRECTORY}/${OUTPUT_GIF_NAME} ${OUTPUT_DIRECTORY}/archives/gif_${d}.gif > /dev/null 2>&1

file_mov=$(ls -t ${MOVIES_DIRECTORY}/*.mov | head -n1)
echo $file_mov

ffmpeg -v ${FFMEPG_LOGLEVEL} -ss 00:00:00.000 -i "${file_mov}" -pix_fmt rgb24 -r 10 -s ${GIF_SIZE} -t 00:00:10.000 ${OUTPUT_DIRECTORY}/output.gif

convert -layers Optimize ${OUTPUT_DIRECTORY}/output.gif ${OUTPUT_DIRECTORY}/${OUTPUT_GIF_NAME} && rm -f ${OUTPUT_DIRECTORY}/output.gif

du -h ${OUTPUT_DIRECTORY}/${OUTPUT_GIF_NAME}