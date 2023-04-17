@echo off
if "%1" == "h" goto begin
mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit
:begin

F:
cd F:\my\software\ios\fileserver

fileserver.exe -p 3011