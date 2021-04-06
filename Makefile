.PHONY: setup
setup:
	aws ec2 import-key-pair --key-name "jenkins-admin-key" --public-key-material fileb://~/.ssh/id_rsa.pub

.PHONY: validate
validate:
	aws cloudformation validate-template --template-body file://vpc.yaml | jq
	aws cloudformation validate-template --template-body file://main.yaml | jq
	aws cloudformation validate-template --template-body file://jenkins.yaml | jq