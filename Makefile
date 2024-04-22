dep:
	go install github.com/gojuno/minimock/v3/cmd/minimock@latest

mock:
	minimock -i github.com/godverv/matreshka-be/pkg/matreshka_api.MatreshkaBeAPIClient -o mocks -g -s "_mock.go"
