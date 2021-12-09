mkdir "Day$1"

cp -R Template/* Day$1
cd "Day$1"
go mod init
wget --load-cookies ../cookies.txt -O in.txt https://adventofcode.com/2021/day/$1/input
python3 -m webbrowser https://adventofcode.com/2021/day/$1
read