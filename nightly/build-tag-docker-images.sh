docker build --force-rm -t "landing_service:nightly_${BUILD_NUMBER}" ../landing_service
docker build --force-rm -t "login_service:nightly_${BUILD_NUMBER}" ../login_service