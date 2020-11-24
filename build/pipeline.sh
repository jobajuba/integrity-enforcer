#!/bin/bash
set -e

echo "IMAGE PUBLISH GOES HERE!"

echo "Pipeline for integrity verifier starting : $(date)"

# Tag images with COMPONENT_TAG_EXTENSION
docker tag ${IV_SERVER_IMAGE_NAME_AND_VERSION} ${REGISTRY}/${IV_IMAGE}:${VERSION}${COMPONENT_TAG_EXTENSION}
docker tag ${IV_LOGGING_IMAGE_NAME_AND_VERSION} ${REGISTRY}/${IV_LOGGING}:${VERSION}${COMPONENT_TAG_EXTENSION}
docker tag ${IV_OPERATOR_IMAGE_NAME_AND_VERSION} ${REGISTRY}/${IV_OPERATOR}:${VERSION}${COMPONENT_TAG_EXTENSION}

export COMPONENT_VERSION=${VERSION}
export COMPONENT_DOCKER_REPO=${REGISTRY}

# Publish ${IV_IMAGE}
export COMPONENT_NAME=${IV_IMAGE}
export DOCKER_IMAGE_AND_TAG=${COMPONENT_DOCKER_REPO}/${COMPONENT_NAME}:${COMPONENT_VERSION}${COMPONENT_TAG_EXTENSION}
rm -rf pipeline
make pipeline-manifest/update PIPELINE_MANIFEST_COMPONENT_SHA256=${TRAVIS_COMMIT} PIPELINE_MANIFEST_COMPONENT_REPO=${TRAVIS_REPO_SLUG} PIPELINE_MANIFEST_BRANCH=${TRAVIS_BRANCH}
echo "Completed pipeline for integrity verifier component: $COMPONENT_NAME"

# Publish ${IV_LOGGING}
export COMPONENT_NAME=${IV_LOGGING}
export DOCKER_IMAGE_AND_TAG=${COMPONENT_DOCKER_REPO}/${COMPONENT_NAME}:${COMPONENT_VERSION}${COMPONENT_TAG_EXTENSION}
rm -rf pipeline
make pipeline-manifest/update PIPELINE_MANIFEST_COMPONENT_SHA256=${TRAVIS_COMMIT} PIPELINE_MANIFEST_COMPONENT_REPO=${TRAVIS_REPO_SLUG} PIPELINE_MANIFEST_BRANCH=${TRAVIS_BRANCH}
echo "Completed pipeline for integrity verifier component: $COMPONENT_NAME"

# Publish ${IV_OPERATOR}
export COMPONENT_NAME=${IV_OPERATOR}
export DOCKER_IMAGE_AND_TAG=${COMPONENT_DOCKER_REPO}/${COMPONENT_NAME}:${COMPONENT_VERSION}${COMPONENT_TAG_EXTENSION}
rm -rf pipeline
make pipeline-manifest/update PIPELINE_MANIFEST_COMPONENT_SHA256=${TRAVIS_COMMIT} PIPELINE_MANIFEST_COMPONENT_REPO=${TRAVIS_REPO_SLUG} PIPELINE_MANIFEST_BRANCH=${TRAVIS_BRANCH}
echo "Completed pipeline for integrity verifier component: $COMPONENT_NAME"

