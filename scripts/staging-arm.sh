#!/bin/bash

export SLUG=ghcr.io/awakari/embed-text-arm64
export VERSION=latest
docker tag awakari/embed-text-arm64 "${SLUG}":"${VERSION}"
docker push "${SLUG}":"${VERSION}"
