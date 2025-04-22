# 使用 Node.js 20 官方镜像
FROM node:20-alpine

# 设置工作目录
WORKDIR /app
# 1. 配置国内镜像源和安装必要工具
RUN npm config set registry https://registry.npmmirror.com && \
    npm install -g pnpm && \
    pnpm config set registry https://registry.npmmirror.com
# 安装 pnpm
RUN npm install -g pnpm

# 复制项目文件
COPY . .

RUN npm install -g tsx
RUN npm install -g express
# 安装依赖
RUN pnpm install

# 暴露端口（如果需要）
EXPOSE 8082

# 启动命令
CMD ["pnpm", "run", "start"]