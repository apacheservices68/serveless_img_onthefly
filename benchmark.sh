#!/bin/bash
#
# You must have installed vegeta:
# go get github.com/tsenart/vegeta
#
port=3300
# Start the server
./bin/tto-resize -origin ./tests/origin -cache ./tests/cache & > /dev/null
pid=$!

suite() {
  echo "$1 --------------------------------------"
  echo "GET http://localhost:$port/$2" | vegeta attack \
    -duration=30s \
    -rate=50 \
    -body="./tests/large.jpg" \ | vegeta report
  sleep 1
}

# Run suites
suite "Crop" "crop/600_400/ttnew/large.jpg"
suite "Zoom" "zoom/600_400/ttnew/large.jpg"
suite "Fit" "fit/300_300/ttnew/large.jpg"
suite "Thumb Width" "thumb_w/200/ttnew/large.jpg"
suite "Thumb Height" "thumb_h/300/ttnew/large.jpg"

# Kill the server
kill -9 $pid