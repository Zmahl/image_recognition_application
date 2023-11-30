package auth

import "os"

func setCredentials(credentials *GoogleCloudCredentials) {
	credentials.VisionApiKey = os.Getenv("VISION_API_KEY")
	credentials.CloudStorageServiceAccount = os.Getenv("GOOGLE_CLOUD_SERVIEC_ACCOUNT")
	credentials.BucketName = os.Getenv("BUCKET_NAME")
}

func NewCloudCredentials() *GoogleCloudCredentials {
	CloudCredentials := &GoogleCloudCredentials{}
	setCredentials(CloudCredentials)
	return CloudCredentials
}
