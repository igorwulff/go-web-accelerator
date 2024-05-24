run:
	@templ generate
	@go:generate npm run build
	@go run ./main.go