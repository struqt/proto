#!/bin/bash
set -euo pipefail

declare -r SELF=$(readlink -f "$0")
declare -r SELF_DIR=${SELF%/*}
declare -r BUF_DIR="${SELF_DIR}"
declare -r GRPC_GATEWAY_VERSION="${1:-v2.26.3}"

download_file_with_curl() {
  local url="https://raw.githubusercontent.com/${1:?}"
  local path="${BUF_DIR:?}/${2:?}"
  mkdir -p "$(dirname "${path:?}")"
  curl "${url:?}" -o "${path}" -s
  echo -e "File downloaded: $url\n ---> $path"
}

download_proto_google_api() {
  local version="${1}"
  download() {
    local file="${1}"
    local repo='googleapis'
    local base_url="googleapis/${repo}/${version:?}/google/api"
    local base_path="${repo}/google/api"
    download_file_with_curl "${base_url}/${file}" "${base_path}/${file:?}"
  }
  download annotations.proto
  download field_behavior.proto
  download http.proto
  download httpbody.proto
}

download_proto_grpc_gateway() {
  local version="${1}"
  download() {
    local file="${1}"
    local package=protoc-gen-openapiv2/options
    local repo='grpc-gateway'
    local base_url="grpc-ecosystem/${repo}/${version:?}/${package}"
    local base_path="${repo}/${package}"
    download_file_with_curl "${base_url}/${file}" "${base_path}/${file:?}"
  }
  download annotations.proto
  download openapiv2.proto
}

download_proto_google_api master
download_proto_grpc_gateway "${GRPC_GATEWAY_VERSION:?}"

echo 'Downloading is finished'
