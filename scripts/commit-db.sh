#! /bin/bash

db_name=sled_db
tstamp=`date +'%Y-%m-%d'`
dpath=$BKPATH/sled
cp build/$db_name $dpath/$db_name-$tstamp
