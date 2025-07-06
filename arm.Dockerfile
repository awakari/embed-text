FROM --platform=linux/arm64 golang:1.24.4-bookworm AS builder
ARG MODEL_TYPE=intfloat
ARG MODEL_NAME=multilingual-e5-small
ARG MODEL_FILE_ONNX=model_O4.onnx
ARG ONNX_RUNTIME_VERSION=1.22.0
ARG TOKENIZER_FILE=tokenizer.json
ARG TOKENIZER_VERSION=1.20.2
WORKDIR /go/src/embed-text
COPY . .
RUN \
    curl -L https://huggingface.co/${MODEL_TYPE}/${MODEL_NAME}/resolve/main/onnx/${MODEL_FILE_ONNX}?download=true -o /model.onnx && \
    curl -L https://huggingface.co/${MODEL_TYPE}/${MODEL_NAME}/resolve/main/${TOKENIZER_FILE}?download=true -o /tokenizer.json && \
    curl -L https://github.com/microsoft/onnxruntime/releases/download/v${ONNX_RUNTIME_VERSION}/onnxruntime-linux-aarch64-${ONNX_RUNTIME_VERSION}.tgz -o /onnxruntime.tgz && \
    tar -xzf /onnxruntime.tgz && \
    cp -f onnxruntime-linux-aarch64-${ONNX_RUNTIME_VERSION}/lib/libonnxruntime.so.${ONNX_RUNTIME_VERSION} /usr/lib/onnxruntime.so && \
    curl -L https://github.com/daulet/tokenizers/releases/download/v${TOKENIZER_VERSION}/libtokenizers.linux-aarch64.tar.gz -o /libtokenizers.tgz && \
    tar -xzf /libtokenizers.tgz && \
    cp -f libtokenizers.a /usr/lib/libtokenizers.a && \
    apt-get update && \
    apt-get install -y \
      protobuf-compiler \
      libprotobuf-dev && \
    make build-arm64

FROM --platform=linux/arm64 debian:bookworm-slim
COPY --from=builder /go/src/embed-text/embed-text /bin/embed-text
COPY --from=builder /model.onnx /model/model.onnx
COPY --from=builder /tokenizer.json /model/tokenizer.json
COPY --from=builder /usr/lib/onnxruntime.so /usr/lib/onnxruntime.so

ENTRYPOINT ["/bin/embed-text"]
