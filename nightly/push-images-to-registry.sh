repo=scottkurt/gallery
images=`docker image ls | grep -i $repo | grep -i nightly | awk '{print $2}'`
docker login docker.io -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}
echo "pushing   images to registry"
for image in $images 
do 
    image=${repo}:${image}
    docker push $image
done 