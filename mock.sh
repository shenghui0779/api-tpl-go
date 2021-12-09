#!/bin/bash
mockgen -destination=pkg/iao/client/mock_client.go -package=client github.com/shenghui0779/yiigo HTTPClient,UploadForm