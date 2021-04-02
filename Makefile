.PHONY: create-stack
create-stack:
	aws cloudformation create-stack --stack-name=MyCompanyVpc --template-body file://vpc.yaml --parameters file://parameters.json

.PHONY: update-stack
update-stack:
	aws cloudformation update-stack --stack-name=MyCompanyVpc --template-body file://vpc.yaml --parameters file://parameters.json

.PHONY: delete-stack
delete-stack:
	aws cloudformation delete-stack --stack-name=MyCompanyVpc

.PHONY: validate
validate:
	aws cloudformation validate-template --template-body file://vpc.yaml | jq