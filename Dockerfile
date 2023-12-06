FROM golang:1.20.4 AS builder

ENV GOOS=js
ENV GOARCH=wasm

COPY wasm/ /go/src/github.com/michaelact/kubernetes-secret-editor/wasm/

RUN cd /go/src/github.com/michaelact/kubernetes-secret-editor/wasm/ \
	&& cp /usr/local/go/misc/wasm/wasm_exec.js wasm_exec.js \
	&& go build -o kubernetes-secret.wasm

FROM nginx:1.25 AS runtime

COPY --from=builder /go/src/github.com/michaelact/kubernetes-secret-editor/wasm/wasm_exec.js /usr/share/nginx/html/
COPY --from=builder /go/src/github.com/michaelact/kubernetes-secret-editor/wasm/kubernetes-secret.wasm /usr/share/nginx/html/
COPY frontend/* /usr/share/nginx/html/

COPY default.conf /etc/nginx/conf.d/default.conf
