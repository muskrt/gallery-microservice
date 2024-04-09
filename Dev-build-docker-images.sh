echo "Running Image Build Script"
repo=scottkurt/gallery
images=`docker image ls | grep -i scottkurt/gallery | awk '{print $2}'`
services=`ls  | grep '_service' `
echo "removing old images"
for image in $images 
do 
    image=${repo}:${image}
    echo $image
    [  -z "$(docker images -q ${image})" ] || docker  image rm $image 
done 
for service in $services 
do 
echo "docker build --force-rm -t "${repo}:${service}_dev_${BUILD_NUMBER}" ./${service}"
docker build --force-rm -t "${repo}:`echo ${service}|tr '_' '-'`_dev_${BUILD_NUMBER}" ./${service}
done


