
#mkdir -p static/u static/temp-images

git clone https://github.com/DanteAlighierin/ftwa.git

cd ftwa

chmod +x qr.sh
chmod +x cert.sh
chmod +x pkg.sh

./pkg.sh
./cert.sh #generate certificates


echo "Dependencies satisfied, certificates generated"
echo "Go to 'ftwa' folder and enter 'go run main.go' to start the service"

rm installer.sh
