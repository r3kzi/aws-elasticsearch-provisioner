# AWS Elasticsearch Service Provisioner

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