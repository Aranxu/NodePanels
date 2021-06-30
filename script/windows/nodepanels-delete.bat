@echo off

set SH_VERSION=v1.0.1
set UPDATE_TIME=2021.06.29

echo *******************************************************************
echo ^|                        __                            __         ^|
echo ^|       ____  ____  ____/ /__  ____  ____ _____  ___  / /____     ^|
echo ^|      / __ \/ __ \/ __  / _ \/ __ \/ __ \/ __ \/ _ \/ / ___/     ^|
echo ^|     / / / / /_/ / /_/ /  __/ /_/ / /_/ / / / /  __/ (__  )      ^|
echo ^|    /_/ /_/\____/\__,_/\___/ .___/\__,_/_/ /_/\___/_/____/       ^|
echo ^|                          /_/                                    ^|
echo ^|                                                                 ^|
echo ^|               delete script version = %SH_VERSION%                    ^|
echo ^|                     update time  = %UPDATE_TIME%                   ^|
echo ^|                                                                 ^|
echo ^|=================================================================^|

::Stop service
if exist "%USERPROFILE%\nodepanels" (

	::Stop service
	echo %date% %time% Stopping probe ......
	net stop Nodepanels-daemon
	net stop Nodepanels-probe
	
	::Uninstall service
	echo %date% %time% Uninstalling probe ......
	sc delete Nodepanels-daemon
	sc delete Nodepanels-probe
	
	::Delete probe file
	rmdir /q /s "%USERPROFILE%\nodepanels" >NUL 2>NUL
)

echo *********************** delete complete ***************************