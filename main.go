package main

import (
	"fmt"
    "os"
    "net"
    "errors"
    //"os/exec"
	"io/ioutil"
	"net/http"
    //"syscall"
    //"github.com/DanteAlighierin/ftwa/ip"
)



func getPort() string{
    p := os.Getenv("PORT")
    if p != ""{
        return ":" + p
    }
    return ":8080"
}





func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	///1. parse input, , type multipart/form-data
	r.ParseMultipartForm(10 << 20)

	//2. retrieve file from posted form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("err")
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Filename)
	fmt.Printf("MIME type: %+v\n", handler.Header)

	//3. write temp file to server
	tempFile, err := ioutil.TempFile("./temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	//4. return whether or not  this has ben succesful
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(getPort(), nil)
}


//just fix



func main() {
	fmt.Println("Go File Upload")
	fmt.Println("The app is running at http://localhost:8080")
    //fmt.Println(ip.externalIP())

//then this piece of code will run the module, and not the shit that I wrote below

    ip, err := internalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)




	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	setupRoutes()

}


//local usage; yes, this is a shitty approach, i know
func internalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
            return "Adress in your local network: http://" + ip.String() + ":8080", nil
		}
	}
	return "", errors.New("are you connected to the network?")
}






