#/bin/bash
rm -rf ./tmp
cp -r terraform tmp

cd tmp/enviroments/dev

digest=$(gcloud artifacts docker images list ${DOCKER_REPOSITORY}/api --format="value(version)" --sort-by ~UPDATE_TIME --limit 1)

sed -i -e "s/{{PROJECT_ID}}/${TF_VAR_project}/g" ./backend.tf
sed -i \
    -e "s|{{DOCKER_REPOSITORY}}|${DOCKER_REPOSITORY}|g" \
    -e "s/{{DIGEST}}/${digest}/g" \
    ./variables.tf
