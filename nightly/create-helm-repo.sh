export PATH=$PATH:/usr/local/bin
AWS_REGION=us-east-1
HELM_S3_MODE=3 
aws s3 rm s3://gallery-helm-charts --recursive
aws s3api create-bucket --bucket gallery-helm-charts --region us-east-1
aws s3api put-object --bucket gallery-helm-charts --key stable/gallery/
helm s3 init s3://gallery-helm-charts/stable/gallery 
aws s3 ls s3://gallery-helm-charts/stable/gallery/
helm repo add stable-galleryapp s3://gallery-helm-charts/stable/gallery/
helm repo ls

##remove repo
#aws s3 rm s3://gallery-helm-charts --recursive