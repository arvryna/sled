#! /bin/bash
# Script to backup DB

db_name=sled_db
tstamp=`date +'%Y-%m-%d'`
dpath=$BKPATH/sled
cp ~/.sled/$db_name $dpath/$db_name-$tstamp
