package computing

import (
	"strings"
	"sync"

	"github.com/filswan/go-mcs-sdk/mcs/api/bucket"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/filswan/go-mcs-sdk/mcs/api/user"
	"github.com/swanchain/go-computing-provider/conf"
)

var storage *StorageService
var storageOnce sync.Once

type StorageService struct {
	McsApiKey      string `json:"mcs_api_key"`
	McsAccessToken string `json:"mcs_access_token"`
	NetWork        string `json:"net_work"`
	BucketName     string `json:"bucket_name"`
	mcsClient      *user.McsClient
}

func NewStorageService() *StorageService {
	storageOnce.Do(func() {
		storage = &StorageService{
			McsApiKey:  conf.GetConfig().MCS.ApiKey,
			NetWork:    conf.GetConfig().MCS.Network,
			BucketName: conf.GetConfig().MCS.BucketName,
		}
		var err error
		var mcsClient *user.McsClient

		if storage.McsAccessToken != "" {
			mcsClient, err = user.LoginByApikey(storage.McsApiKey, storage.McsAccessToken, storage.NetWork)
		} else {
			mcsClient, err = user.LoginByApikeyV2(storage.McsApiKey, storage.NetWork)
		}

		if err != nil {
			logs.GetLogger().Errorf("Failed creating mcsClient, error: %v", err)
			return
		}
		storage.mcsClient = mcsClient
	})

	return storage
}

func (storage *StorageService) UploadFileToBucket(objectName, filePath string, replace bool) (*bucket.OssFile, error) {
	logs.GetLogger().Infof("uploading file to bucket, objectName: %s, filePath: %s", objectName, filePath)
	buketClient := bucket.GetBucketClient(*storage.mcsClient)

	file, err := buketClient.GetFile(storage.BucketName, objectName)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		logs.GetLogger().Errorf("Failed get file form bucket, error: %v", err)
		return nil, err
	}

	if file != nil {
		if err = buketClient.DeleteFile(storage.BucketName, objectName); err != nil {
			logs.GetLogger().Errorf("Failed delete file form bucket, error: %v", err)
			return nil, err
		}
	}

	if err := buketClient.UploadFile(storage.BucketName, objectName, filePath, replace); err != nil {
		logs.GetLogger().Errorf("Failed upload file to bucket, error: %v", err)
		return nil, err
	}

	mcsOssFile, err := buketClient.GetFile(storage.BucketName, objectName)
	if err != nil {
		logs.GetLogger().Errorf("Failed get file form bucket, error: %v", err)
		return nil, err
	}
	return mcsOssFile, nil
}

func (storage *StorageService) DeleteBucket(bucketName string) error {
	return bucket.GetBucketClient(*storage.mcsClient).DeleteBucket(bucketName)
}

func (storage *StorageService) CreateBucket(bucketName string) {
	_, err := bucket.GetBucketClient(*storage.mcsClient).CreateBucket(bucketName)
	if err != nil {
		logs.GetLogger().Errorf("Failed create bucket, error: %v", err)
		return
	}
}

func (storage *StorageService) CreateFolder(folderName string) {
	_, err := bucket.GetBucketClient(*storage.mcsClient).CreateFolder(storage.BucketName, folderName, "")
	if err != nil {
		logs.GetLogger().Errorf("Failed create folder, error: %v", err)
		return
	}
}

func (storage *StorageService) GetGatewayUrl() (*string, error) {
	return bucket.GetBucketClient(*storage.mcsClient).GetGateway()
}
