#Set variable
PROBE_PATH=/usr/local/nodepanels
SH_VERSION=v1.0.1

echo "*******************************************************************"
echo "|                        __                            __         |"
echo "|       ____  ____  ____/ /__  ____  ____ _____  ___  / /____     |"
echo "|      / __ \/ __ \/ __  / _ \/ __ \/ __ \/ __ \/ _ \/ / ___/     |"
echo "|     / / / / /_/ / /_/ /  __/ /_/ / /_/ / / / /  __/ (__  )      |"
echo "|    /_/ /_/\____/\__,_/\___/ .___/\__,_/_/ /_/\___/_/____/       |"
echo "|                          /_/                                    |"
echo "|                                                                 |"
echo "|               delete script version = ${SH_VERSION}                    |"
echo "|                                                                 |"
echo "|=================================================================|"

#Stop service
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

#Delete probe
if [ -d "${PROBE_PATH}" ]
then
	rm -Rf ${PROBE_PATH}
fi

#Delete service
if [ -f "/etc/systemd/system/nodepanels.service" ]
then
	rm -Rf /etc/systemd/system/nodepanels.service
	rm -Rf /etc/systemd/system/nodepanels-daemon.service
	systemctl daemon-reload
fi

#Delete script
rm -rf $0

echo ""
echo "*********************** delete success ***************************"
