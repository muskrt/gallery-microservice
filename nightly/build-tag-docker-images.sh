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
services=`ls  | grep '_service'`
for service in $services 
do 
docker build --force-rm -t "${repo}:${service}_${build_env}_${BUILD_NUMBER}" ./${service}
done 