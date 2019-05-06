@ECHO OFF

@REM copy config
xcopy /Y .\config\*.* .\bin\

@IF %ERRORLEVEL% NEQ 0 PAUSE