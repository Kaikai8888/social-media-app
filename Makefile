.PHONY: docker
docker:	
	# get first argument

	# 刪除上次編譯的東西 (|| true 代表如果刪除失敗(e.g. 沒有webook檔案), 不會造成整個指令失敗)
	@rm webook || true
	@docker rmi -f kaijump/webook:v0.0.1 || true

	# 執行go mod tidy, 避免go.sum文件不對, 造成編譯失敗
	@go mod tidy

	# 指定編譯成在ARM架構的linux操作系統上運行的可執行文件
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook

	# 建立docker image
	@docker build -t kaijump/webook:v0.0.1 .
