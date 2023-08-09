#!/bin/bash

PKGDIR=watchyourlan-$1-0_all

umask 0022

mkdir -p $PKGDIR/usr/bin
mkdir -p $PKGDIR/lib/systemd/system

cp configs/watchyourlan.service $PKGDIR/lib/systemd/system/

cp watchyourlan $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: watchyourlan
Version: $1
Section: utils
Priority: optional
Depends: arp-scan, tzdata
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: Lightweight network IP scanner with web GUI
" > $PKGDIR/DEBIAN/control

echo "
systemctl daemon-reload
" > $PKGDIR/DEBIAN/postinst

chmod 775 $PKGDIR/DEBIAN/postinst

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR