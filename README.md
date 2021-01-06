# ftwa


![](https://heroku-badge.herokuapp.com/?app=ftwa)
![](https://img.shields.io/github/languages/count/DanteAlighierin/ftwa)
![](https://img.shields.io/github/languages/top/DanteAlighierin/ftwa)
![](https://img.shields.io/github/repo-size/DanteAlighierin/ftwa)
![](https://img.shields.io/github/issues-pr/DanteAlighierin/ftwa)
![](https://img.shields.io/github/issues-pr-closed/DanteAlighierin/ftwa)
![](https://img.shields.io/github/last-commit/DanteAlighierin/ftwa)
![](https://img.shields.io/github/contributors/DanteAlighierin/ftwa)
![](https://img.shields.io/github/license/DanteAlighierin/ftwa)



file transfer web application


### Desktop

![](screens/light_d.png)
![](screens/dark_d.png)

### Mobile

<img src="screens/light_m.png" height="540px" width="270px"><img src="screens/dark_m.png" height="540px" width="270px">






## Public Instances


[Main](https://ftwa.herokuapp.com) - demo showing applied UI/UX solutions

[Testing](https://secure-lake-20134.herokuapp.com) - technodemo to show how the application works


## Core features

- [x] simple and readable code

- [x] access to your files from anywhere in the world or from the local
  network

  - [ ] git synchronizaation

  - [x] encryption

  - [x] privacy and confidentiality

  - [x] convenient and nice looking interface

    - [x] file preview(now only in upload interface)

    - [x] dark theme 

    - [x] adaptive design



## Getting the source code

### via git clone

```bash

git clone https://github.com/DanteAlighierin/ftwa.git
cd ftwa

```


### via go get

```bash

go get github.com/DanteAlighierin/ftwa
cd go/src/github.com/DanteAlighierin/ftwa

```


## Installation

Satisfy dependencies and generate certificates manually:

```bash

##generate certificates
sudo chmod +x cert.sh
./cert.sh

```
or use the script:

```bash
sudo chmod +x installer.sh
./installer.sh

```


Run app interactively:

```bash

go run main.go

```

or build and install the binary:

```bash
go build
go install
```

## Issues & bugs

- unreadable filenames of uploaded files

- loss of cerificates
