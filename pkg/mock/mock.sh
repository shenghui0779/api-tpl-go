#!/bin/bash
mockgen -source=../dao/user.go -destination=dao_user.go -package=mock
mockgen -source=../iao/user/client.go -destination=iao_user.go -package=mock