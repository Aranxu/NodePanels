#!/bin/bash

#Set variable
PROBE_PATH=/usr/local/nodepanels

PROBE_NAME=nodepanels-probe
DAEMON_NAME=nodepanels-daemon

PROBE_SERVICE_NAME=nodepanels
DAEMON_SERVICE_NAME=nodepanels-daemon

SH_VERSION=v1.0.4
PROBE_VERSION=v1.0.2
UPDATE_TIME=2021.12.01

echo "*******************************************************************"
echo "|                        __                            __         |"
echo "|       ____  ____  ____/ /__  ____  ____ _____  ___  / /____     |"
echo "|      / __ \/ __ \/ __  / _ \/ __ \/ __ \/ __ \/ _ \/ / ___/     |"
echo "|     / / / / /_/ / /_/ /  __/ /_/ / /_/ / / / /  __/ (__  )      |"
echo "|    /_/ /_/\____/\__,_/\___/ .___/\__,_/_/ /_/\___/_/____/       |"
echo "|                          /_/                                    |"
echo "|                                                                 |"
echo "|                   script version = ${SH_VERSION}                       |"
echo "|                    probe version = ${PROBE_VERSION}                       |"
echo "|                     update time  = ${UPDATE_TIME}                   |"
echo "|                                                                 |"
echo "|=================================================================|"

#Parameters check
if [ $# -lt 1 ]
then
	echo "|    Can't find probe id\n"
	exit 1
fi

########################################################################################################

#Stop service step 1
HASSERVICE=0
service nodepanels status > /dev/null 2>&1
if [ $? != 4 ]
then
	echo "$(date +"%Y-%m-%d %T") Stopping probe ......"
	service ${DAEMON_SERVICE_NAME} stop
	service ${PROBE_SERVICE_NAME} stop
	HASSERVICE=1
fi

#Stop service step 2
PROCESS=`ps -ef|grep nodepanels-daemon|grep -v grep|grep -v PPID|awk '{ print $2}'`
for i in $PROCESS
do
  kill -9 $i
done

PROCESS=`ps -ef|grep nodepanels-probe|grep -v grep|grep -v PPID|awk '{ print $2}'`
for i in $PROCESS
do
  kill -9 $i
done
	
#Delete service
if [ ${HASSERVICE} = 1 ]
then

	echo "$(date +"%Y-%m-%d %T") Removing probe service ......"
	
	if [ -d ${PROBE_PATH} ]
	then
		${PROBE_PATH}/${DAEMON_NAME} -uninstall
		${PROBE_PATH}/${PROBE_NAME} -uninstall
	fi
	
	systemctl disable ${DAEMON_SERVICE_NAME}
	systemctl disable ${PROBE_SERVICE_NAME}
	systemctl daemon-reload	
	
fi

if [ -f "/etc/systemd/system/nodepanels.service" ]
then
	echo "$(date +"%Y-%m-%d %T") Removing probe service ......"
	rm -Rf /etc/systemd/system/nodepanels.service
	rm -Rf /etc/systemd/system/nodepanels-daemon.service
	systemctl daemon-reload
fi
	
#Delete previous probe path
if [ -d ${PROBE_PATH} ]
then
	echo "$(date +"%Y-%m-%d %T") Deleting nodepanels path ......"
	rm -Rf ${PROBE_PATH}
fi

########################################################################################################

#Create probe dir
echo "$(date +"%Y-%m-%d %T") ========== Start to install nodepanels-probe =========="
echo "$(date +"%Y-%m-%d %T") Creating nodepanels path ......"
mkdir -p ${PROBE_PATH}

#Judge arch set download url
PROBE_DOWNLOAD_URL=""
DAEMON_DOWNLOAD_URL=""
if [ `uname -m` = "i686" ] ; then
	PROBE_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/nodepanels-probe-linux-386"
	DAEMON_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/nodepanels-daemon-linux-386"
elif [ `uname -m` = "x86_64" ] ; then
	PROBE_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/nodepanels-probe-linux-amd64"
	DAEMON_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/nodepanels-daemon-linux-amd64"
elif [ `uname -m` = "aarch32" ] ; then
	PROBE_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/nodepanels-probe-linux-arm32"
	DAEMON_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/nodepanels-daemon-linux-arm32"
elif [ `uname -m` = "aarch64" ] ; then
	PROBE_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/prod/nodepanels-probe-linux-arm64"
	DAEMON_DOWNLOAD_URL="https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/daemon/prod/nodepanels-daemon-linux-arm64"
else
	echo "$(date +"%Y-%m-%d %T") The system is not supported for the time being"
	exit
fi

#Download probe
echo "$(date +"%Y-%m-%d %T") Downloading probe ......"
wget -nv -o /dev/stdout -O ${PROBE_PATH}/${PROBE_NAME} --no-check-certificate ${PROBE_DOWNLOAD_URL}
echo "$(date +"%Y-%m-%d %T") Download probe success"

#Download daemon
echo "$(date +"%Y-%m-%d %T") Downloading daemon ......"
wget -nv -o /dev/stdout -O ${PROBE_PATH}/${DAEMON_NAME} --no-check-certificate ${DAEMON_DOWNLOAD_URL}
echo "$(date +"%Y-%m-%d %T") Download daemon success"

#Grant authority
chmod +x ${PROBE_PATH}/${PROBE_NAME}
echo "$(date +"%Y-%m-%d %T") Grant probe success"
chmod +x ${PROBE_PATH}/${DAEMON_NAME}
echo "$(date +"%Y-%m-%d %T") Grant daemon success"

#Create config file
echo "{\"serverId\":\""$1"\"}" >> ${PROBE_PATH}/config
echo "$(date +"%Y-%m-%d %T") Create config file success"

#Register as a service
${PROBE_PATH}/${PROBE_NAME} -install
${PROBE_PATH}/${DAEMON_NAME} -install

service nodepanels-daemon restart

echo "$(date +"%Y-%m-%d %T") Install success"

#Delete script
rm -rf $0

echo "*********************** install success ***************************"
