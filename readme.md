# Currency exchange calculator  

Currency exchange calculator written in GO, uses data from [NBP API](http://api.nbp.pl/). Current version converts values between PLN and GBP.  
App has gui and command line version:
* gui version available on [main branch](https://github.com/SzymekN/currency-exchange-calculator)
* Command line version available on [no_gui_branch](https://github.com/SzymekN/currency-exchange-calculator/tree/no_gui_version)

## Table of contents
- [Currency exchange calculator](#currency-exchange-calculator)
  - [Table of contents](#table-of-contents)
  - [Technologies](#technologies)
  - [How to run](#how-to-run)
    - [Non gui version](#non-gui-version)
    - [Gui version](#gui-version)
  - [How to use](#how-to-use)
    - [gui version](#gui-version-1)
    - [Non gui version](#non-gui-version-1)


## Technologies  
* [GO v1.20](https://go.dev/) - main programming language
  * [Go-app](https://go-app.dev/) - package used to build progressive web apps (PWA)
  * [gomock](https://github.com/golang/mock) - mocking framework used in testing
* HTML5 and CSS3 - creating and styling web gui 
## How to run  
In order to use either of available versions you have to download and install [Go](https://go.dev/dl/) in version 1.20 or newer.

### Non gui version

To run the program all you have to do is to download the repo or clone it. To run the program you can use **main.exe** located in main directory or if you want to rebuild and run the program you can execute **go run .** in command line opened in main directory of the program.

### Gui version

First step is the same - download the repo or clone it. Now you can used compiled binary **main.exe** to run the server and then you can access the website through your web browser on address http://localhost:8000/.

If you want to make changes in code and run it again you have to first build **app.wasm** artifact (static client) using either of the commands below:
* **GOARCH=wasm GOOS=js go build -o web/app.wasm** - on Linux
* **cmd /v /c "set GOOS=js&& set GOARCH=wasm&& go build -o web/app.wasm"** - on Windows

After that you can build server using **go build -o main.exe**. After that you can run the program as described at the beginning of this section (running main.exe and accessing it through http://localhost:8000/).

In main directory there is also script named **build.bat**, that builds client and server artifacts - **this script works only on Windows.**

## How to use
### gui version 
After accessing the page you should see website like this below. Typing in either of input fields will be detected and calculated values (send/received) will be displayed in other input field.  
![image](https://user-images.githubusercontent.com/83112762/221240612-fa9da276-7d74-4a2b-944d-c7a67855987f.png)

Send(GBP) to receive(PLN)  
![image](https://user-images.githubusercontent.com/83112762/221242209-a03df23a-a4a4-4135-893d-466df2e46e08.png)

Receive(PLN) to send(GBP)
![image](https://user-images.githubusercontent.com/83112762/221241777-67b5c424-13ec-48e9-bbe9-4a3087b95cbc.png)

### Non gui version
After running main.exe you should see window like this below. Then you can choose one of the listed option:
* 1 - calculate from GBP to PLN (send -> receive)
* 2 - calculate from PLN to GBP (receive -> send)

![image](https://user-images.githubusercontent.com/83112762/221244365-7b18147f-8384-49f4-b48b-5894329f72c9.png)

After choosing one option you can enter amount and calculated value will be printed, then you can again choose what do you want to calculate.  
![image](https://user-images.githubusercontent.com/83112762/221244646-ac09833a-54cb-49f9-911b-d1edda7ff8ee.png)
![image](https://user-images.githubusercontent.com/83112762/221244750-adb5a840-c165-4712-9f8e-7707e0cc3a62.png)

