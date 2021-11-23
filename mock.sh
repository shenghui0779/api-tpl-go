#!/bin/bash
mockgen -source=pkg/dao/user.go -destination=pkg/dao/mock_user.go -package=dao