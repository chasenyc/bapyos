## BAPYOS

![in use](/assets/screenshot.png)
## Installation
```sh
git clone https://github.com/chasenyc/bapyos
```
then cd into directory
```sh
go build bapyos.go
```
and finally:
```sh
./bapyos
```

## Setup
In order for this app to work properly, we need three folders:
1. ba
1. pyos
3. combined
Within ba and pyos you need to have corresponding files with the same name. So aladdin.svg in `/ba` and aladdin.svg in `/pyos`. When you run the binary all you will need to input is `aladdin` and it will find the `.svg` file in both directories and output the result to the `/combined` folder as `aladdin.svg`.

see below image to illustrate folder structure:
![folder structure](/assets/screenshot2.png)