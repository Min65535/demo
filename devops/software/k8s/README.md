## deploy-shell
```text
#!/usr/bin/env bash

set -e

## env
# shellcheck disable=SC2034
ROOT=$1
kubectl version
if [[ "${CI_COMMIT_REF_NAME}" == "dev" ]]; then export RUN_ENV=dev; fi
if [[ "${CI_COMMIT_REF_NAME}" == "master" ]]; then export RUN_ENV=test; fi
case $RUN_ENV in
"dev") ;;
"test") ;;
*)
  echo "ci_deploy env=dev|test"
  exit 0
  ;;
esac
echo "RUN_ENV: ""${RUN_ENV}"

## get image name
COMMIT=${CI_COMMIT_SHA}
BUILD_DATE=$(date +%Y%m%d)
VERSION='v1.0.0'
VERSION=${VERSION}-${COMMIT}-build-${BUILD_DATE}
IMAGE_TAG=my-services:${VERSION}
DOMAIN_NAME="docker-registry.my.net"
IMAGE=${DOMAIN_NAME}/my-${RUN_ENV}/${IMAGE_TAG}
if [ "$RUN_ENV" == "dev" ]; then
  IMAGE=${DOMAIN_NAME}/my-test/${IMAGE_TAG}
fi
echo "IMAGE_TAG: ""${IMAGE_TAG}"
echo "IMAGE: ""${IMAGE}"

## generate k8s deploy dir
DirBinK8s=${ROOT}/bin/k8s
#DirBinK8sHash=${ROOT}/bin/k8s/${RUN_ENV}_${CI_COMMIT_SHA}
DirBinK8sHash=${DirBinK8s}/${RUN_ENV}_${COMMIT}
DirK8sDeploy=${ROOT}/deploy/kubernetes
FileK8sDeploy=${DirK8sDeploy}/deploy.yml
mkdir -p "${DirBinK8s}"
# shellcheck disable=SC2115
rm -rf "${DirBinK8s}"/*
mkdir -p "${DirBinK8sHash}"
cd "${DirBinK8sHash}"

## get name in k8s svc discovery
function SvcGetName() {
  svcLatest=""
  case $1 in
  "one-service")
    # shellcheck disable=SC2034
    svcLatest="one-svr"
    ;;
  "two-service")
    # shellcheck disable=SC2034
    svcLatest="two-svr"
    ;;
  *)
    echo "bad service to get name"
    exit 0
    ;;
  esac
}

## get port in k8s svc discovery
function SvcGetPort() {
  port=0
  case $1 in
  "one-service")
    # shellcheck disable=SC2034
    port=2000
    ;;
  "two-service")
    # shellcheck disable=SC2034
    port=2001
    ;;
  *)
    echo "bad service to get port"
    exit 0
    ;;
  esac
}

## get apps
# shellcheck disable=SC2012
apps=$(ls -lh "${ROOT}"/app | awk 'NR>1{print $9}')
echo "${apps}"

for i in ${apps}; do
  echo "--------------++++++++++++++++----------------"
  echo "start to generate svc deployment file: ""${i}"
  depFile="deploy-""${i}"".yml"
  cp -r "${FileK8sDeploy}" "${DirBinK8sHash}"/"${depFile}"
  sed -i "s/_run_env_var_/${RUN_ENV}/g" "${depFile}"
  sed -i "s/_app_name_var_/${i}/g" "${depFile}"
  sed -i "s/_namespace_var_/my-test-${RUN_ENV}/g" "${depFile}"
  sed -i "s! _image_name_var_! ${IMAGE}! g" "${depFile}"
  SvcGetName "${i}"
  name=$svcLatest
  sed -i "s/_svc_var_/${name}/g" "${depFile}"
  SvcGetPort "${i}"
  # shellcheck disable=SC2006
  # shellcheck disable=SC2003
  nodePort=$(expr $port + 30000)
  sed -i "s/_port_var_/${port}/g" "${depFile}"
  sed -i "s/_nodePort_var_/${nodePort}/g" "${depFile}"
  echo "end generating svc deployment file: ""${i}"
  echo "--------------==================----------------"
  cat "${depFile}"

  echo "start to apply svc deployment file: ""${i}"
  kubectl apply -f "${depFile}"
  echo "successfully apply svc deployment file: ""${i}"
done

exit 0
```