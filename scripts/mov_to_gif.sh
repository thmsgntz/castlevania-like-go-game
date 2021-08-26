
# https://superuser.com/questions/436056/how-can-i-get-ffmpeg-to-convert-a-mov-to-a-gif
# https://stackoverflow.com/questions/34341808/is-there-a-way-to-add-a-gif-to-a-markdown-file

set +x

ffmpeg -ss 00:00:00.000 -i $1 -pix_fmt rgb24 -r 10 -s 640x480 -t 00:00:10.000 ressources/gifs/output.gif
convert -layers Optimize ressources/gifs/output.gif ressources/gifs/output_optimized.gif

du -h ressources/gifs/output_optimized.gif