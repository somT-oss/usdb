package main

import (
	"fmt"
	"log"
	"os"
	"usdb/storage"
)

func main() {
	var maxPage uint32 = 4
	var bufferedPages uint32 = 0
	var pageContent [4096]byte
	filePath := "wal.log"
	file, err := os.OpenFile("wal.log", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("failed to open file in read write mode.")
		panic(err)
	}

	wal, err := storage.NewWAL(
		maxPage,
		bufferedPages,
		pageContent,
		file,
		filePath,
	)

	rowInformation := `
1,Somto,DevOps Engineer,AWS,Kubernetes,Lagos,Nigeria
2,Ada,Backend Engineer,Go,PostgreSQL,Abuja,Nigeria
3,John,Cloud Engineer,Terraform,Docker,Nairobi,Kenya
4,Mary,SRE,Prometheus,Grafana,Cape Town,South Africa
5,David,Platform Engineer,Kafka,Kubernetes,Accra,Ghana
6,Jane,Security Engineer,IAM,CloudTrail,Kigali,Rwanda
7,Michael,Data Engineer,Spark,Airflow,Cairo,Egypt
8,Sarah,Software Engineer,Python,Django,Lagos,Nigeria
9,Daniel,Infra Engineer,Ansible,Ubuntu,Douala,Cameroon
10,Grace,DevOps Engineer,EKS,Helm,Dakar,Senegal
11,Chris,Cloud Architect,VPC,IAM,Johannesburg,South Africa
12,Ruth,SRE,Linux,eBPF,Lagos,Nigeria
13,Peter,Backend Engineer,Redis,NATS,Kampala,Uganda
14,Evelyn,DevOps Engineer,CI/CD,GitHub Actions,Lagos,Nigeria
15,Victor,Infrastructure Engineer,NGINX,HAProxy,Abuja,Nigeria
16,Angela,Platform Engineer,ArgoCD,Kustomize,Nairobi,Kenya
17,Samuel,Cloud Engineer,S3,Lambda,Accra,Ghana
18,Diana,Database Engineer,PostgreSQL,WAL,Cairo,Egypt
19,Henry,SRE,Kubernetes,etcd,Lagos,Nigeria
20,Joy,Backend Engineer,Rust,Tonic,Douala,Cameroon
21,Frank,DevOps Engineer,Docker,Compose,Kigali,Rwanda
22,Alice,Infrastructure Engineer,Terraform,AWS,Dakar,Senegal
23,Brian,Cloud Engineer,GCP,Kubernetes,Johannesburg,South Africa
24,Nancy,Security Engineer,Vault,Secrets,Nairobi,Kenya
25,Emmanuel,Platform Engineer,Helm,Ingress,Lagos,Nigeria
26,Chinedu,SRE,Alertmanager,Prometheus,Abuja,Nigeria
27,Patricia,Backend Engineer,Go,gRPC,Cape Town,South Africa
28,Kevin,DevOps Engineer,ECS,Fargate,Accra,Ghana
29,Linda,Infra Engineer,BGP,Networking,Cairo,Egypt
30,Joseph,Cloud Architect,Multi-Cloud,Kubernetes,Lagos,Nigeria
31,Rebecca,Database Engineer,MySQL,Replication,Douala,Cameroon
32,Isaac,Platform Engineer,GitOps,FluxCD,Nairobi,Kenya
33,Hannah,SRE,Chaos Engineering,Lagos,Nigeria
34,Paul,Backend Engineer,Microservices,RabbitMQ,Kampala,Uganda
35,Faith,Cloud Engineer,CloudFront,S3,Abuja,Nigeria
36,Daniela,DevOps Engineer,Loki,Grafana,Dakar,Senegal
37,Stephen,Infra Engineer,Systemd,Linux,Cape Town,South Africa
38,Naomi,Security Engineer,OIDC,OAuth2,Lagos,Nigeria
39,Elijah,Platform Engineer,Service Mesh,Istio,Accra,Ghana
40,Sophia,SRE,Monitoring,Observability,Nairobi,Kenya
41,Matthew,Cloud Engineer,EC2,AutoScaling,Cairo,Egypt
42,Olivia,Backend Engineer,NodeJS,MongoDB,Lagos,Nigeria
43,Benjamin,DevOps Engineer,GitLab CI,ArgoCD,Johannesburg,South Africa
44,Victoria,Infra Engineer,ZFS,Storage,Douala,Cameroon
45,Andrew,Database Engineer,SQLite,BTrees,Kampala,Uganda
46,Deborah,SRE,Incident Response,PagerDuty,Abuja,Nigeria
47,Joshua,Cloud Architect,Hybrid Cloud,AWS,Dakar,Senegal
48,Melissa,Platform Engineer,OpenTelemetry,Tracing,Lagos,Nigeria
49,Anthony,Backend Engineer,CQRS,Event Sourcing,Nairobi,Kenya
50,Clara,Security Engineer,SIEM,Wazuh,Cape Town,South Africa
`
	if err != nil {
		panic(err)
	}
	err = wal.Push(1, []byte(rowInformation))
	if err != nil {
		panic(err)
	}
	fmt.Println("Done writing row to wal.log")
}
