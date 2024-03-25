docker build --force-rm -t "landing_service:dev_${BUILD_NUMBER}" ./landing_service
docker build --force-rm -t "login_service:dev_${BUILD_NUMBER}" ./login_service