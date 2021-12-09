#!/bin/bash
mockgen -destination=pkg/mock/mock_http.go -package=mock github.com/shenghui0779/yiigo HTTPClient,UploadForm
