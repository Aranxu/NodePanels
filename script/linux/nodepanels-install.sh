#!/bin/bash

#Set variable
PROBE_PATH=/usr/local/nodepanels
PROBE_NAME=nodepanels-probe
DAEMON_NAME=nodepanels-daemon
SH_VERSION=v1.0.2
PROBE_VERSION=v1.0.1
UPDATE_TIME=2021.06.20

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

#Delete previous probe
if [ -d ${PROBE_PATH} ]
then
	echo "$(date +"%Y-%m-%d %T") Delete historical version"
	rm -Rf ${PROBE_PATH}
	
	#Stop service
	echo "$(date +"%Y-%m-%d %T") Stopping probe ......"
	service nodepanels-daemon stop
	service nodepanels stop
	systemctl disable nodepanels-daemon
	systemctl disable nodepanels
	
	PROCESS=`ps -ef|grep nodepanels-probe|grep -v grep|grep -v PPID|awk '{ print $2}'`
	for i in $PROCESS
	do
	  kill -9 $i
	done
	
	PROCESS=`ps -ef|grep nodepanels-daemon|grep -v grep|grep -v PPID|awk '{ print $2}'`
	for i in $PROCESS
	do
	  kill -9 $i
	done
	
	#Delete service
	if [ -f "/etc/systemd/system/nodepanels.service" ]
	then
		echo "$(date +"%Y-%m-%d %T") Removing probe service ......"
		rm -Rf /etc/systemd/system/nodepanels.service
		rm -Rf /etc/systemd/system/nodepanels-daemon.service
		systemctl daemon-reload
	fi

fi

#Create probe dir
mkdir -p ${PROBE_PATH}

#Download probe
echo "$(date +"%Y-%m-%d %T") Downloading probe ......"

if [ `arch` = "aarch64" ] ; then
   wget -nv -o /dev/stdout -O ${PROBE_PATH}/${PROBE_NAME} --no-check-certificate https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/arm/nodepanels-probe
else
   wget -nv -o /dev/stdout -O ${PROBE_PATH}/${PROBE_NAME} --no-check-certificate https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/x86/nodepanels-probe
fi

echo "$(date +"%Y-%m-%d %T") Download probe success"

#Download daemon
echo "$(date +"%Y-%m-%d %T") Downloading daemon ......"

if [ `arch` = "aarch64" ] ; then
   wget -nv -o /dev/stdout -O ${PROBE_PATH}/${DAEMON_NAME} --no-check-certificate https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/arm/nodepanels-daemon
else
   wget -nv -o /dev/stdout -O ${PROBE_PATH}/${DAEMON_NAME} --no-check-certificate https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/probe/x86/nodepanels-daemon
fi

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
cat > /etc/systemd/system/nodepanels.service <<EOF
[Unit]
Description=nodepanels

[Service]
ExecStart=/usr/local/nodepanels/nodepanels-probe

[Install]
WantedBy=multi-user.target
EOF

echo "$(date +"%Y-%m-%d %T") Register probe service success"

cat > /etc/systemd/system/nodepanels-daemon.service <<EOF
[Unit]
Description=nodepanels-daemon

[Service]
ExecStart=/usr/local/nodepanels/nodepanels-daemon

[Install]
WantedBy=multi-user.target
EOF

echo "$(date +"%Y-%m-%d %T") Register daemon service success"

echo "$(date +"%Y-%m-%d %T") Starting probe ......"
systemctl enable nodepanels
systemctl enable nodepanels-daemon
service nodepanels-daemon restart

#Delete script
rm -rf $0

echo "*********************** install success ***************************"
