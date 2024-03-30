echo "Running Image Build Script For Nightly Build"
build_env=nightly
repo=scottkurt/gallery
images=`docker image ls | grep -i scottkurt/gallery | awk '{print $2}'`
services=`docker image ls | awk '{print $1}' | grep service `
echo "removing old images"
for image in $images 
do 
    image=${repo}:${image}
    echo $image
    [  -z "$(docker images -q ${image})" ] || docker  image rm $image 
done 
docker build --force-rm -t "${repo}:landing_service_${build_env}_${BUILD_NUMBER}" ../landing_service
docker build --force-rm -t "${repo}:login_service_${build_env}_${BUILD_NUMBER}" ../login_service
images=`docker image ls | grep -i scottkurt/gallery | awk '{print $2}'`
services=`docker image ls | awk '{print $1}' | grep service `
echo "removing  images"
for image in $images 
do 
    image=${repo}:${image}
    echo $image
    [  -z "$(docker images -q ${image})" ] || docker  image rm $image 
done 