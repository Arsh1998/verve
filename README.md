# verve
Technical challenge solution by Arshdeep Singh

# running service docker image
`docker run -p 8080:8080 \
    -e AWS_ACCESS_KEY_ID=<your-access-key> \
    -e AWS_SECRET_ACCESS_KEY=<your-secret-key> \
    -e AWS_REGION=<your-region> \
    verve-service`