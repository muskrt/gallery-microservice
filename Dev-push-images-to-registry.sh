repo=scottkurt/gallery
images=`docker image ls | grep -i scottkurt/gallery | awk '{print $2}'`
USERNAME=$1
PASSWORD=$2
docker login docker.io -u $USERNAME -p $PASSWORD
echo "pushing   images to registry"
for image in $images 
do 
    image=${repo}:${image}
    docker push $image
done 


