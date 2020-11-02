#!/usr/bin/bash


cat /etc/*-release


if grep -q void /etc/*-release

then sudo xbps-install go libressl

fi


if grep -q slackware /etc/*-release

then sudo slackpkg install gcc-go openssl

fi
