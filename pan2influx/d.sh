#!/usr/bin/env bash
#sudo setenforce 0
docker stop grafana-xxl
docker rm grafana-xxl
#docker run -d -p 3003:3003 --name grafana -h grafana grafana/grafana
#docker cp grafana:/etc/grafana/grafana.ini .
#docker run -d -p 3003:3003 --name grafana --volume=/home/docker/grafana:/var/lib/grafana -v /home/docker/grafana.ini:/etc/grafana/grafana.ini:ro -h grafana grafana/grafana

# create /var/lib/grafana as persistent volume storage
#docker run -d -v /var/lib/grafana --name grafana-xxl-storage busybox:latest

# start grafana-xxl
docker run \
  -d --network influx \
  -p 3003:3003 \
  --name grafana-xxl \
  --volumes-from grafana-xxl-storage -v $PWD/grafana/defaults.ini:/usr/share/grafana/conf/defaults.ini:ro \
  -v $PWD/grafana/server.key:/usr/share/grafana/conf/server.key -v $PWD/grafana/grafana.cer:/usr/share/grafana/conf/grafana.cer \
  -v $PWD/grafana/ldap.toml:/etc/grafana/ldap.toml  monitoringartist/grafana-xxl:latest

#docker run --rm influxdb influxd config > influxdb.conf
docker stop influx
docker rm influx
docker run -d --network influx -p 8083:8083 -p 8086:8086 --name influx -h influx\
      -v $PWD/influxdb:/var/lib/influxdb -v $PWD/influxdb.conf:/etc/influxdb/influxdb.conf:ro\
      influxdb -config /etc/influxdb/influxdb.conf