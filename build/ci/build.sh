#!/bin/bash

cd cmd/WatchYourLAN/ && CGO_ENABLED=0 go build -o ../../watchyourlan .
cd ../../

PKGDIR=/opt/watchyourlan

umask 0022

mkdir -p $PKGDIR
cp watchyourlan $PKGDIR/
cp configs/watchyourlan.service $PKGDIR/
cp configs/install.sh $PKGDIR/

cd /opt
tar cvzf watchyourlan-$1.tar.gz watchyourlan
cd -
cp /opt/watchyourlan-$1.tar.gz .