# Base Setup

Basic docker-compose setup. It helps to place and monitor several non-critical services\prototipes on one machine. 
This setup contains:
1. [Traefic](https://doc.traefik.io/traefik/)
2. [Fluentd](https://www.fluentd.org/)
3. [Prometheus](https://prometheus.io/) + [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/) (with telegram notification)
4. [Mongo](https://www.mongodb.com/) as a log storage

Apps are writing logs in JSON format, and fluentd parses these logs writes them to mongo, and pushes metrics to Prometheus.
An example format of a log format can be found in a "testservice" folder.

## Testing the example
To start this example you have to create alermanager config `conf.yaml` in the folder `alertmanager/config` by replacing `ALERT_BOT_TOKEN` and `CHAT_ID` in `conf_template.yaml`. You can get the token by this [instruction](https://core.telegram.org/bots/api#authorizing-your-bot) from official telegram documentation and as a `CHAT_ID` you can use your chat ID from `@RawDataBot`.

You can start the example with this command. 
```
docker compose --env-file=dev.env up --force-recreate --build
```
Click this link http://a.localhost:8080/abc to get a message about "invalid id". This request produces a log message in one of the services. 
```json
{"code":0,"message":"Failed to parse strconv.Atoi: parsing \"abc\": invalid syntax","level":"error","handler":"default"}
```
Connect to mongo `fluentd` base with credentials from `dev.env`. You'll see that for every service will be created collection with log records. 