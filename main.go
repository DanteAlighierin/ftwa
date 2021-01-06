package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/ccding/go-stun/stun"
)

func KeyHandler() {

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	func() {

		<-c

		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
		number := GoRoutineChannel()
		close(number)

	}()

}

func GoRoutineChannel() chan int {

	ch := make(chan int)

	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

// get port name
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	} //exception
	return httpsAddr
}

func generator() {
	genTree := exec.Command("./downloads.sh")

	genOut, err := genTree.Output()

	if err != nil {

		panic(err)

	}

	fmt.Println(string(genOut))
}

//main func of program.

func uploadFile(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Uploading File\n")

	///1. parse input, , type multipart/form-data
	req.ParseMultipartForm(10 << 20)

	//2. retrieve file from posted form data
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("err")
		return
	}
	defer file.Close() //print file's data
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v", handler.Size)
	fmt.Printf(" KB\n")
	fmt.Printf("MIME type: %+v\n", handler.Header)

	//3. write temp file to server
	tempFile, err := ioutil.TempFile("./static/temp-images", "upload-*")
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
	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	http.ServeFile(w, req, "./static/upload")

	generator()
	//5. gen downloads page via tree util

	//genTree := exec.Command("downloads.sh")

	//genOut, err := genTree.Output()

	//if err != nil {
	//  panic(err)
	//}
	//fmt.Println(string(genOut))
}

var httpAddr string = ":8080"
var httpsAddr string = ":8443"

// upload routing
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	//http.ListenAndServe(getPort(), nil)
	//forcing use of TLS(https protocol)
	//http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
	//http.ListenAndServe(":8080", http.HandlerFunc(redir))

	srv := http.Server{
		Addr: httpsAddr,
	}
	_, tlsPort, err := net.SplitHostPort(httpsAddr)

	if err != nil {
		return
	}

	go redirectToHTTPS(tlsPort)
	srv.ListenAndServeTLS("server.crt", "server.key")

	//go http.ListenAndServe(":8080", http.HandlerFunc(redir))
}

func redirectToHTTPS(tlsPort string) {
	httpSrv := http.Server{
		Addr: httpAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, _ := net.SplitHostPort(r.Host)
			u := r.URL
			u.Host = net.JoinHostPort(host, tlsPort)
			u.Scheme = "https"
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}

//greetings
func main() {
	fmt.Println("Welcome to ftwa - file transfer web application.")
	fmt.Println("The app is running at https://localhost" + getPort())
	//fmt.Println(ip.externalIP())
	//KeyHandler()
	//then this piece of code will run the module, and not the shit that I wrote below

	testCmd := exec.Command("./qr.sh")

	testOut, err := testCmd.Output()

	if err != nil {

		panic(err)
	}
	generator()
	//fmt.Println(string(testOut))

	//print public ip via using the module

	ip, err := internalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	ipex, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Public address: https://" + ipex + getPort())
	fmt.Println(string(testOut))
	//number := GoRoutineChannel()
	//fmt.Println(<-number)
	//fmt.Println(<-number)
	//close(number)
	go KeyHandler()

	// set served directory
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
			return "Address in your local network: http://" + ip.String() + getPort(), nil
		}
	}
	return "", errors.New("are you connected to the network?")

}

//get external ip via OVERBLOATED FUCKING WEBSITE. Bad practice
func ExternalIP() (string, error) {

	rsp, err := http.Get("https://myexternalip.com/raw")

	if err != nil {

		return "", err

	}

	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)

	if err != nil {

		return "", err

	}

	//return data

	return string(bytes.TrimSpace(buf)), nil

}

// go get "github.com/ccding/go-stun/stun"
func NewExternalIP() (string, error) {
	client := stun.NewClient()
	client.SetServerAddr(stun.DefaultServerAddr)
	client.SetVerbose(true)  //for debug
	client.SetVVerbose(true) //for debug
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("NAT Type:", nat)
	if host != nil {
		fmt.Println("External IP Family:", host.Family())
		fmt.Println("External IP:", host.IP())
		fmt.Println("External Port:", host.Port())
		return host.IP(), nil
	}
	return "", errors.New("Some error")
}
