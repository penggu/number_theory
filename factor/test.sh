go build
for i in `seq 100 1000`; do ./factor $i; done
