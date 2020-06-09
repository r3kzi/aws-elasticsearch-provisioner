<p align="left"><img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/8b1d3e63f2013bf48b04c906312dc358f6f916e3.png" alt="Logo" height="200"></p> 

# AWS Elasticsearch Service Provisioner

[![Go Report Card](https://goreportcard.com/badge/github.com/r3kzi/aws-elasticsearch-provisioner)](https://goreportcard.com/report/github.com/r3kzi/aws-elasticsearch-provisioner)

You want to provision your AWS Elasticsearch Service Cluster and you're using an IAM Master User with 
[Fine-Grained Access Control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html)?

You can't use Ansible because it doesn't allow you to sign your HTTP requests with 
[AWS Signature V4](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)?

Fear no more! I've built something!

## Configuration

| Parameter                 | Description                                                                               | Default                                           |
|---------------------------|-------------------------------------------------------------------------------------------|---------------------------------------------------|
| `elasticsearch.endpoint`  | Configurable [AWS Elasticsearch Service](https://aws.amazon.com/de/elasticsearch-service) | `https://elasticsearch`                           |
| `aws.region`              | AWS Region where your Domain was placed                                                   | `eu-west-1`                                       |
| `aws.roleARN`             | IAM Master User ARN that you defined within Fine-Grained-Access-Control settings          | `arn:aws:iam::123456123456:role/IAMMasterUser`    |

## Installation

AWS Elasticsearch Service Provisioner requires a
[supported release of Go](https://golang.org/doc/devel/release.html#policy).

```shell script
$ go get -u github.com/r3kzi/elasticsearch-provisioner
```

To find out where `golint` was installed you can run `$ go list -f {{.Target}} github.com/r3kzi/elasticsearch-provisioner`. 

For `elasticsearch-provisioner` to be used globally add that directory to the `$PATH` environment setting.

## Contributing

Pull requests are welcome.