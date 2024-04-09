echo "Running Image Build Script For Nightly Build"
build_env=nightly
repo=scottkurt/gallery
images=`docker image ls | grep -i scottkurt/gallery | grep -i nightly | awk '{print $2}'`
echo "removing old images"
for image in $images 
do 
    image=${repo}:${image}
    echo $image
    [  -z "$(docker images -q ${image})" ] || docker  image rm $image 
done 


docker build --force-rm -t "${IMAGE_TAG_APIGW_SERVICE}" ./apigw_service
docker build --force-rm -t "${IMAGE_TAG_LOGIN_SERVICE}" ./login_service
docker build --force-rm -t "${IMAGE_TAG_LANDING_SERVICE}" ./landing_service
docker build --force-rm -t "${IMAGE_TAG_NOTES_SERVICE}" ./notes_service
