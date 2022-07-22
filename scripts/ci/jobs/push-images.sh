#!/usr/bin/env bash

# Adapted from https://github.com/stackrox/stackrox/blob/master/scripts/ci/jobs/push-images.sh

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")"/../../.. && pwd)"
source "$ROOT/scripts/ci/lib.sh"

set -euo pipefail

push_images() {
    info "Will push images built in CI"

    info "Images from OpenShift CI builds:"
    env | grep IMAGE || true

    [[ "${OPENSHIFT_CI:-false}" == "true" ]] || { die "Only supported in OpenShift CI"; }

    push_image_set

    if is_in_PR_context; then
        comment_on_pr
    fi
}

comment_on_pr() {
    info "Adding a comment with the build tag to the PR"

    # hub-comment is tied to Circle CI env
    local url
    url=$(get_pr_details | jq -r '.html_url')
    export CIRCLE_PULL_REQUEST="$url"

    local sha
    sha=$(get_pr_details | jq -r '.head.sha')
    sha=${sha:0:7}
    export _SHA="$sha"

    local tag
    tag=$(make tag)
    export _TAG="$tag"

    local tmpfile
    tmpfile=$(mktemp)
    cat > "$tmpfile" <<- EOT
Images are ready for the commit at {{.Env._SHA}}.

To use with deploy scripts, first \`export MAIN_IMAGE_TAG={{.Env._TAG}}\`.
EOT

    hub-comment -type build -template-file "$tmpfile"
}

push_images "$@"
