#!/usr/bin/bash


cat /etc/*-release #find out the distro


if grep -q void /etc/*-release #Search by keyword
then sudo xbps-install go libressl #Installing packages using the package manager of a specific distro
fi


if grep -q slackware /etc/*-release #Search by keyword
then sudo slackpkg install gcc-go openssl #Installing packages using the package manager of a specific distro
fi


if grep -q nixos /etc/*-release #Search by keyword
then nix-env -i gcc-go openssl #Installing packages using the package manager of a specific distro
fi


if grep -q "ubuntu | mint | debian" /etc/*-release #Search by keyword
then sudo apt install golang openssl #Installing packages using the package manager of a specific distro
fi



if grep -q "arch | manjaro" /etc/*-release #Search by keyword
then sudo pacman -S golang openssl #Installing packages using the package manager of a specific distro
fi



./cert.sh #generate certificates


echo "Dependencies satisfied, certificates generated"
echo "Enter 'go run main.go' to start the service"

