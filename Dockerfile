FROM ghcr.io/knights-analytics/hugot:v0.3.3 AS builder
ARG MODEL_TYPE=sentence-transformers
ARG MODEL_NAME=paraphrase-multilingual-MiniLM-L12-v2
ARG MODEL_FILE_ONNX=model_O3.onnx
ARG TOKENIZER_FILE=tokenizer.json
WORKDIR /go/src/embed-text
COPY . .
RUN \
    curl -L https://huggingface.co/${MODEL_TYPE}/${MODEL_NAME}/resolve/main/onnx/${MODEL_FILE_ONNX}?download=true -o /model.onnx && \
    curl -L https://huggingface.co/${MODEL_TYPE}/${MODEL_NAME}/resolve/main/${TOKENIZER_FILE}?download=true -o /tokenizer.json && \
    dnf install -y \
      protobuf-compiler \
      protobuf-devel && \
    make build

FROM ghcr.io/knights-analytics/hugot:v0.3.3 AS runtime
COPY --from=builder /go/src/embed-text/embed-text /bin/embed-text
COPY --from=builder /model.onnx /model/model.onnx
COPY --from=builder /tokenizer.json /model/tokenizer.json

ENTRYPOINT ["/bin/embed-text"]
