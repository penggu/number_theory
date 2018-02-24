go build

start=2
stop=10
if [[ ! -z $1 ]]; then start=$1; fi
if [[ ! -z $2 ]]; then stop=$2; fi

for i in `seq $start $stop`; do ./factor $i; done
