.PHONY:

HUB_NAME = dentalchart.kr.ncr.ntruss.com

protoc_gen_go:
	@read -p "Enter Proto File Name:" filename; \
	./third_party/protoc/bin/protoc.exe $$filename --go_out=. ;