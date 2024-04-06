#!/usr/bin/env bash


if [ `whoami` != root ]; then
    echo Please run this script as root or using sudo
    exit
fi

set -e
cd `dirname $0`


chown -R elder:elder /opt/elder
chmod -R 0600 /opt/elder
chmod -R u+rwX /opt/elder
chmod u+x /opt/elder/app
