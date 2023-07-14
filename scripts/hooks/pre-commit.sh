#!/bin/bash

# check_golangci_lint_version ensures that the local version of golangci-lint
# matches the version specified in .github/workflows/golangci.yml. This makes sure
# that if the pre-commit hook is run locally, the changes should align with CI.
function check_golangci_lint_version(){
  local git_root="$(git rev-parse --show-toplevel)"

  # Note: we are explicitly stripping out the 'v' prefix from the versions. Different verisons of
  # golangci-lint have different version formats. For example, v1.27.0 (if installed with go get) vs 1.27.0 (installed with curl).

  # extract the version of golangci-lint from the CI workflow file.
  local workflow_golangci_lint_version="$(grep ' version' ${git_root}/.github/workflows/golangci.yml | awk '{ print $NF }' | sed "s/v//g" )"

  local local_golangci_lint_version="$(golangci-lint version --format short | grep '[0-9\.]'| sed "s/v//g")"

  if [[ "${workflow_golangci_lint_version}" != "${local_golangci_lint_version}" ]]; then
    echo "local golangci-lint (${local_golangci_lint_version}) must be upgraded to ${workflow_golangci_lint_version}"
    echo "aborting pre-commit hook"
    exit 1
  fi
}

# run_hook formats all modified go files that have been staged and adds them to the commit.
function run_hook() {
  make lint-fix-changed
  echo "formatting any staged go files"
  go_files="$(git diff --staged --name-only | grep \.go$)"
  for f in $go_files; do
    git add $f
  done
}

check_golangci_lint_version
run_hook