#!/bin/bash

export SLUG=ghcr.io/awakari/embed-text
export VERSION=latest
docker tag awakari/embed-text "${SLUG}":"${VERSION}"
docker push "${SLUG}":"${VERSION}"
