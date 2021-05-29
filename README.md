# practice-elk(docker-compose)

The practice create simple and speed to usage ELK.

## pre-prepare 
- install docker/docker-compose

## Usage

### start
```
bash start.sh
```
Wait a few minutes... use link in the below  

Link  
[kibana](http://127.0.0.1:5601)  
[elasticsearch](http://127.0.0.1:9200)  
[nginx](http://127.0.0.1:8080)  

### stop
```
bash stop.sh 
```

## project folder structure

├── README.md  
├── docker-compose.yml  
├── elasticsearch.yml  
├── filebeat.yml  
├── kibana.yml  
├── logstash.conf  
├── logstash.yml  
└── var  


## Configuration
Configuration have five files. `elasticsearch.yml`, `filebeat.yml`, `kibana.yaml`, `logstash.conf`, `logstash.yml`

