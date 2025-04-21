#!/bin/bash

# 一键安装Golang环境并运行项目
# 使用方法：chmod +x setup_golang.sh && ./setup_golang.sh

set -e  # 遇到错误立即退出

# 定义变量
GO_VERSION="1.23.5"
INSTALL_DIR="/usr/local"
PROJECT_DIR="$(pwd)"

# 检查是否为root用户
if [ "$(id -u)" -ne 0 ]; then
  echo "请使用root用户或通过sudo运行此脚本"
  exit 1
fi

# 安装依赖
echo "安装系统依赖..."
apt-get update -qq
apt-get install -y -qq wget tar git

# 下载并安装Golang
echo "安装Go ${GO_VERSION}..."
if [ -d "${INSTALL_DIR}/go" ]; then
  echo "检测到已安装的Go环境，先进行清理..."
  rm -rf "${INSTALL_DIR}/go"
fi

wget -qO- "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | \
tar -C "${INSTALL_DIR}" -xzf -

# 设置环境变量
echo "配置环境变量..."
cat >> /etc/profile <<EOF
export GOROOT=${INSTALL_DIR}/go
export GOPATH=\$HOME/go
export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH
EOF

source /etc/profile

# 验证安装
echo "验证安装..."
go version || { echo "Go安装失败"; exit 1; }

# 配置Go模块代理
echo "设置Go模块代理..."
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 进入项目目录
cd "${PROJECT_DIR}"

# 下载依赖
echo "执行 go mod tidy..."
go mod tidy || { echo "go mod tidy 失败"; exit 1; }

# 运行项目
echo "启动项目..."
go run . || { echo "go run 失败"; exit 1; }

echo "✅ 所有操作已完成"