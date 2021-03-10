#!/usr/bin/bash


chmod +x qr.sh
chmod +x cert.sh

if grep -q void /etc/*-release #Search by keyword
then sudo xbps-install go libressl qrencode #Installing packages using the package manager of a specific distro
fi


if grep -q slackware /etc/*-release #Search by keyword
then sudo slackpkg install gcc-go openssl qrencode #Installing packages using the package manager of a specific distro
fi


if grep -q nixos /etc/*-release #Search by keyword
then nix-env -i gcc-go openssl #Installing packages using the package manager of a specific distro
fi


if grep -q 'ubuntu\|debian\|mint' /etc/*-release #Search by keyword
then sudo apt install golang openssl qrencode #Installing packages using the package manager of a specific distro
fi


if grep -q 'arch\|manjaro' /etc/*-release #Search by keyword
then sudo pacman -S go openssl qrencode #Installing packages using the package manager of a specific distro
fi


#mkdir -p static/u static/temp-images



./cert.sh #generate certificates


echo "Dependencies satisfied, certificates generated"
echo "Enter 'go run main.go' to start the service"

