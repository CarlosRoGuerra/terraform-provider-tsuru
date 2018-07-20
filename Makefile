build:
	go build -o terraform-provider-tsuru

tfinit:
	terraform init

tfplan:
	terraform plan

tfapply:
	terraform apply
