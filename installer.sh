./pkg.sh

#mkdir -p static/u static/temp-images

git clone https://github.com/DanteAlighierin/ftwa.git

cd ftwa

rm installer.sh

chmod +x qr.sh
chmod +x cert.sh

./cert.sh #generate certificates


echo "Dependencies satisfied, certificates generated"
echo "Enter 'go run main.go' to start the service"

