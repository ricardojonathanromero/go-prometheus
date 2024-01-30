# go-prometheus
This is an example of prometheus and a simple golang app

## Dependencies
* Golang ([see version][golang-version] recommended)
* Prometheus ([see version][prometheus-version] recommended)

## Explanation
The service health only configures a health service and the metrics service.
The first is used to return a health checker, and the second is used by prometheus recovery some information of the server.

The Prometheus server used a certificate (cer and key files) that allows to configure the TLS feature but at this moment
this feature is in mode experimental and could change in the future, so, please check the [documentation](https://prometheus.io/docs/prometheus/latest/configuration/https/)
if you want to modify the configuration.

Other consideration is the basic auth config. This feature is enabled and is configured into the `configs/web-config.yml:11`
The user is `john` and the password is `Temporal1*`, those values are necessary if you want to access to the web console.

Finally, the password mentioned prev is also configured in the `prometheus.yml:27` as a text file called `prometheus_psw.txt`
the password is used to call the endpoint `/metrics` of the prometheus server.

## Login
Go to the url `https://localhost:9090` and you will see a page like this:

![login_page](assets/login.svg)

Use this credentials to access:
* **username:** `john`
* **password:** `Temporal1*`

⚠️ If the browser shows an alert `This page is unsafe`, skip the alert and continue to login page.

## Alerts
The alerts are configured using the `prometheus-rules.yml` file. This contains the rules and the alerts will be configured
on the server. If you desire add a new alert or rule, check the next files: [alerts](docs/prometheus.example.alerts.yml) and
[rules](docs/prometheus.example.rules.yml); where you will find the information about the yml file and some examples.

By default, this project configures two alerts and one rule.

To see the alerts, go to this [link](https://localhost:9090/alerts?search=), and you'll show this page:

![alerts_page](assets/alerts.svg)

The alerts configured monitoring when an instance has been down by more than 5 minutes, and when a http service has a high latency.


To see the rules, go to this [link](https://localhost:9090/rules).

![rules_page](assets/rules.svg)

The rules configured will be described here.


[golang-version]: .goversion
[prometheus-version]: .prometheusversion
