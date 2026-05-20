package storage

import (
	"os"
	"strings"
	"testing"
)

const (
	maxPage uint32 = 4
	bufferedPages uint32 = 0
	filePath = "../wal.log"
	rowInformation = `
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
)

var pageContent [4096]byte

func TestIncreaseBufferedPagesCapacity(t *testing.T) {
	// the current maxPages you can write to the buffer is 4
	// i want to test this by writing more than 4 and confirming the guard works for this implementation.
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		t.Log("failed to open wal file...")
		panic(err)
	}
	wal, err := NewWAL(
		maxPage,
		bufferedPages,
		pageContent,
		file,
		filePath,
	)

	if err != nil {
		t.Logf("failed to initialize wal struct...")
		panic(err)
	}

	for i := range maxPage + uint32(1) {
		t.Logf("writing %v bytes into the buffer...", i * uint32(4096))
		err = wal.Push(i, []byte(rowInformation))
		if err != nil {
			t.Logf("maxPages limit hit at index %v. could only write %v bytes into the buffer", i, i * uint32(4096))
		}
	}
}

func TestReadFromBuffer(t *testing.T) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		t.Log("failed to open wal file...")
		panic(err)
	}
	wal, err := NewWAL(
		maxPage,
		bufferedPages,
		pageContent,
		file,
		filePath,
	)

	if err != nil {
		t.Log("failed to initialize wal struct...")
		panic(err)
	}
	cleanedRowInformation := strings.TrimSpace(rowInformation)

	// push 4096 bytes into the wal buffer.
	wal.Push(1, []byte(cleanedRowInformation))
	
	rowToRead := "1,Somto,DevOps Engineer,AWS,Kubernetes,Lagos,Nigeria"
	byteOffset := len([]byte(rowToRead)) 
	t.Logf("the first row has a length of %d", byteOffset)
	
	firstRow := make([]byte, byteOffset + 4)
	_, err = wal.buffer.Read(firstRow)

	if err != nil {
		t.Errorf("failed to read first row in the buffer...")
		panic(err)
	}

	// we're reading from index 4: because [0:3] contains the 4bytes representing the page number.
	if string(firstRow)[4:] != rowToRead {
		t.Errorf("failed to read the first row in the buffer properly...")
	}
}


