
sh 'Dev-build-packages.sh'
sh 'Dev-generate-coverage.sh'
sh 'Dev-build-docker-images.sh'
sh 'Dev-push-images-to-registry.sh'
sh 'nightly/build-tag-docker-images.sh'