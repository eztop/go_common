all: genenv generrcode
genenv: cleanenv
	@protoc  --go_out=env env/env.proto
	@mv env/github.com/eztop/go_common/env/env.pb.go env
	@rm -rf env/github.com
generrcode: cleanerrcode
	@protoc  --go_out=errcode errcode/errcode.proto
	@mv errcode/github.com/eztop/go_common/errcode/errcode.pb.go errcode
	@rm -rf errcode/github.com
cleanenv:
	@rm -rf env/env.pb.go
cleanerrcode:
	@rm -rf errcode/errcode.pb.go