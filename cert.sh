#!/bin/bash

openssl genrsa -out server.key 2048

openssl ecparam -genkey -name secp384r1 -out server.key

while true; do
	read -p "Do you wanna to autogenerate certs? " yn
	case $yn in


#select yn in "Yes" "No"
#	case $yn in




		[Yy]* ) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/C=EN/ST= /L= /O= /CN=localhost"; echo "Certs generated!!!"; break;;
		

		[Nn]* ) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650;;
		* ) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/C=EN/ST= /L= /O= /CN=localhost"; echo "Certs generated!!!"; break;;

	esac
done





