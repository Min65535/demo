## windows
```text
命令行win+r,输入shell:startup ,再把.bat文件放入打开的文件夹下(此为系统开机自启动程序的文件夹)
```

## shell
```text
@echo off
if "%1" == "h" goto begin
mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit
:begin

F:
cd F:\my\software\ios\fileserver

fileserver.exe -p 3011
```