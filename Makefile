.PHONY: setup
setup:
	aws ec2 import-key-pair --key-name "jenkins-admin-key" --public-key-material fileb://~/.ssh/id_rsa.pub

.PHONY: upload
upload:
	go run s3upload.go

.PHONY: create-stack
create-stack:
	aws cloudformation create-stack --stack-name=MyCompanyVpc --template-body file://main.yaml --parameters file://parameters.json

.PHONY: update-stack
update-stack:
	aws cloudformation update-stack --stack-name=MyCompanyVpc --template-body file://main.yaml --parameters file://parameters.json

.PHONY: delete-stack
delete-stack:
	aws cloudformation delete-stack --stack-name=MyCompanyVpc

.PHONY: validate
validate:
	aws cloudformation validate-template --template-body file://vpc.yaml | jq
	aws cloudformation validate-template --template-body file://main.yaml | jq
	aws cloudformation validate-template --template-body file://jenkins.yaml | jq