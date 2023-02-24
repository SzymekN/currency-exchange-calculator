# Currency exchange calculator  

Currency exchange calculator written in GO, uses data from [NBP API](http://api.nbp.pl/). Current version converts values between PLN and GBP.  
App has gui and command line version:
* gui version available on [main branch](https://github.com/SzymekN/currency-exchange-calculator)
* Command line version available on [no_gui_branch](https://github.com/SzymekN/currency-exchange-calculator/tree/no_gui_version)

## Technologies  
* [GO v1.20](https://go.dev/) - main programming language
  * [Go-app](https://go-app.dev/) - package used to build progressive web apps (PWA)
  * [gomock](https://github.com/golang/mock) - mocking framework used in testing
* HTML5 and CSS3 - creating and styling web gui 
## How to use  
In order to use either of available versions you have to download and install [Go](https://go.dev/dl/) in version 1.20 or newer.

### Non gui version

To run the program all you have to do is to download the repo or clone it. To run the program you can use **main.exe** located in main directory or if you want to rebuild and run the program you can execute **go run .** in command line opened in main directory of the program.

### Gui version

First step is the same - download the repo or clone it. Now you can used compiled binary **main.exe** to run the server and then you can access the website through your web browser on address http://localhost:8000/.

If you want to make changes in code and run it again you have to first build **app.wasm** artifact (static client) using either of the commands below:
* **GOARCH=wasm GOOS=js go build -o web/app.wasm** - on Linux
* **cmd /v /c "set GOOS=js&& set GOARCH=wasm&& go build -o web/app.wasm"** - on Windows

After that you can build server using **go build -o main.exe**. After that you can run the program as described at the beginning of this section (runnin main.exe and accessing it through http://localhost:8000/).
