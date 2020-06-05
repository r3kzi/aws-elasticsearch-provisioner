# AWS Elasticsearch Service Provisioner

You want to provision your AWS Elasticsearch Service Cluster and you're using an IAM Master User with 
[Fine-Grained Access Control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html)?

You can't use Ansible because it doesn't allow you to sign your HTTP requests with 
[AWS Signature V4](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)?

Fear no more! I've built something!