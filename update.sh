#!/bin/bash
set -eo pipefail

declare -rx BUF_DIR='.'

download_protobuf_with_curl() {
  local url="https://raw.githubusercontent.com/${1:?}"
  local path="${BUF_DIR:?}/${2:?}"
  if [ -f "${path:?}" ]; then
    echo "File exists: $path"
    return
  fi
  mkdir -p "$(dirname "${path:?}")"
  curl "${url:?}" -o "${path}"
}

download_protobuf_google_api() {
  local version="${1}"
  download() {
    local file="${1}"
    local repo='googleapis'
    local base_url="googleapis/${repo}/${version:?}/google/api"
    local base_path="${repo}/google/api"
    download_protobuf_with_curl "${base_url}/${file}" "${base_path}/${file:?}"
  }
  download annotations.proto
  download field_behavior.proto
  download http.proto
  download httpbody.proto
}

download_protobuf_grpc_gateway() {
  local version="${1}"
  download() {
    local file="${1}"
    local package=protoc-gen-openapiv2/options
    local repo='grpc-gateway'
    local base_url="grpc-ecosystem/${repo}/${version:?}/${package}"
    local base_path="${repo}/${package}"
    download_protobuf_with_curl "${base_url}/${file}" "${base_path}/${file:?}"
  }
  download annotations.proto
  download openapiv2.proto
}

download_protobuf_google_api master
download_protobuf_grpc_gateway v2.16.0
echo 'Downloading is finished'
