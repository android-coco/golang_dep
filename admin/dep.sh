#!/bin/bash

exe_path="$( cd "$( dirname "$0"  )" && pwd  )"
DEP_HOME=${exe_path}/../
DEP_SBIN=${DEP_HOME}/bin/
DEP_PID=${DEP_HOME}data/dep.pid


start()
{
    if [ -f ${DEP_PID} ]
    then
        pid=`cat ${DEP_PID}`
        process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "dep" | wc -l`
        if [ ${process_num} -ge 1 ];
        then
            echo "service already running. pid=$pid"
            return
        fi  
    fi
    cd ${DEP_SBIN}
    nohup ./dep &> ../logs/dep.log 2>> ../logs/dep_except.log &
    echo "dep start"
}


stop()
{
    if [ ! -f ${DEP_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${DEP_PID}`
    process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "dep" | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    kill -TERM `cat ${DEP_PID}`
    ret=$?
    if [ ${ret} -eq 0 ]
    then
        echo "dep stop"
    else
        echo "dep stop failed"
    fi
    return
}

restart()
{
    stop
    start
    return
}


reload()
{
    if [ ! -f ${DEP_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${DEP_PID}`
    process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "dep" | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    kill -USR2 `cat ${DEP_PID}`
    return
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    reload)
        reload
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart|reload}"
        exit 1
esac

exit 0
