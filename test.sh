#!/bin/bash -u
orig=docker-cloner-test
copy=${orig}-copy
memory_limit=200
expected_memory_bytes=$(( $memory_limit * 1024 * 1024 ))

docker rm -f $orig >/dev/null 2>&1
docker rm -f $copy >/dev/null 2>&1

set -e

docker run --rm -d --name $orig alpine sleep 120 >/dev/null
container_id=$(go run main.go \
-memory 200 \
-id $orig \
-cmd "echo In $copy contianer. PASSED" \
-new_name $copy)

actual_memory_bytes=$(docker inspect $container_id|command jq -r '.[0].HostConfig.Memory')
if [ $actual_memory_bytes -ne $expected_memory_bytes ];then
  echo "Memory limit mismatch: expected $expected_memory_bytes. got $actual_memory_bytes"
  exit 1
fi

docker start -ia $container_id

# cleanup
docker rm -f $orig >/dev/null 2>&1
