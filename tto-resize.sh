#!/bin/bash
case $1 in
    start)
        echo "Starting tto-resize."
        /root/go/bin/tto-resize -origin ./tests/ttnew -cache ./tests/ttnew / &
        ;;
    stop)
        echo "Stopping tto-resize."
        sudo kill $(sudo lsof -t -i:3300)
        ;;
    restart)
        echo "Restarting tto-resize."
        sudo kill $(sudo lsof -t -i:3300)
        /root/go/bin/tto-resize -origin ./tests/ttnew -cache ./tests/ttnew / &
        ;;
    *)
        echo "tto-resize service."
        echo $"Usage $0 {start|stop|restart}"
        exit 1
esac
exit 0