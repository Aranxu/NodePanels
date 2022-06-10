@echo off

set SH_VERSION=v1.0.4
set PROBE_VERSION=v1.1.0
set DAEMON_VERSION=v1.1.0
set UPDATE_TIME=2022.06.10

echo *******************************************************************
echo ^|                        __                            __         ^|
echo ^|       ____  ____  ____/ /__  ____  ____ _____  ___  / /____     ^|
echo ^|      / __ \/ __ \/ __  / _ \/ __ \/ __ \/ __ \/ _ \/ / ___/     ^|
echo ^|     / / / / /_/ / /_/ /  __/ /_/ / /_/ / / / /  __/ (__  )      ^|
echo ^|    /_/ /_/\____/\__,_/\___/ .___/\__,_/_/ /_/\___/_/____/       ^|
echo ^|                          /_/                                    ^|
echo ^|                                                                 ^|
echo ^|                   script version = %SH_VERSION%                       ^|
echo ^|                    probe version = %PROBE_VERSION%                       ^|
echo ^|                     update time  = %UPDATE_TIME%                   ^|
echo ^|                                                                 ^|
echo ^|=================================================================^|

::Administrator check
net.exe session 1>NUL 2>NUL || (
    echo Please run as Administrator
	exit /b 1
)

::Parameters check
if [%1] == [] (
	echo Can't find probe id\n
	exit /b 1
)

::Delete previous probe
if exist "%USERPROFILE%\nodepanels" (

	::Stop service
	echo %date% %time% Stopping probe ......
	sc stop Nodepanels-daemon
	sc stop Nodepanels-probe
	
	::Uninstall service
	echo %date% %time% Uninstalling probe ......
	sc delete Nodepanels-daemon
	sc delete Nodepanels-probe
	
	::Delete probe file
	rmdir /q /s "%USERPROFILE%\nodepanels" >NUL 2>NUL
)

::Create probe dir
md %USERPROFILE%\nodepanels 

::Download probe
echo %date% %time% Downloading probe ......

echo %PROCESSOR_ARCHITECTURE% | find /i "x86" > nul
if %errorlevel%==0 (
	set PROBE_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/%PROBE_VERSION%/nodepanels-probe-windows-386.exe'
	set DAEMON_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/%DAEMON_VERSION%/nodepanels-daemon-windows-386.exe'
)
echo %PROCESSOR_ARCHITECTURE% | find /i "AMD64" > nul
if %errorlevel%==0 (
	set PROBE_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/%PROBE_VERSION%/nodepanels-probe-windows-amd64.exe'
	set DAEMON_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/%DAEMON_VERSION%/nodepanels-daemon-windows-amd64.exe'
)
echo %PROCESSOR_ARCHITECTURE% | find /i "ARM64" > nul
if %errorlevel%==0 (
	set PROBE_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/%PROBE_VERSION%/nodepanels-probe-windows-arm64.exe'
	set DAEMON_DOWNLOAD_URL='https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/%DAEMON_VERSION%/nodepanels-daemon-windows-arm64.exe'
)

powershell -Command "$wc = New-Object System.Net.WebClient; $wc.DownloadFile(%PROBE_DOWNLOAD_URL%, '%USERPROFILE%\nodepanels\nodepanels-probe.exe')"
if errorlevel 1 (
  echo %date% %time% Download Probe Failed
  exit /b 1
)
echo %date% %time% Download probe success

::Download daemon
echo %date% %time% Downloading daemon ......

powershell -Command "$wc = New-Object System.Net.WebClient; $wc.DownloadFile(%DAEMON_DOWNLOAD_URL%, '%USERPROFILE%\nodepanels\nodepanels-daemon.exe')"
if errorlevel 1 (
  echo %date% %time% Download Daemon Failed
  exit /b 1
)
echo %date% %time% Download daemon success

::Create config file
echo {"serverId":"%1"} >> %USERPROFILE%\nodepanels\config.json
echo %date% %time% Create config file success

::Register as a service
%USERPROFILE%\nodepanels\nodepanels-probe.exe -install
echo %date% %time% Register probe service success

%USERPROFILE%\nodepanels\nodepanels-daemon.exe -install
echo %date% %time% Register daemon service success

echo %date% %time% Starting probe ......
net start Nodepanels-daemon

echo *********************** install complete ***************************
