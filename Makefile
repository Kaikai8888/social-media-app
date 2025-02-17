.PHONY: docker
docker:	
	# get first argument

	# 刪除上次編譯的東西 (|| true 代表如果刪除失敗(e.g. 沒有social-media-app檔案), 不會造成整個指令失敗)
	@rm social-media-app || true
	@docker rmi -f kaijump/social-media-app:v0.0.1 || true

	# 執行go mod tidy, 避免go.sum文件不對, 造成編譯失敗
	@go mod tidy

	# 指定編譯成在ARM架構的linux操作系統上運行的可執行文件
	@GOOS=linux GOARCH=arm go build -tags=k8s -o social-media-app

	# 建立docker image
	@docker build -t kaijump/social-media-app:v0.0.1 .
wire_test: 
	@cd ./internal/integration_test/startup; wire