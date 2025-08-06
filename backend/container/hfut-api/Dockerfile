# 使用阿里云官方提供的 Node.js 20 镜像（非 library 命名空间）
FROM node:20-slim

# 后续内容保持不变...
WORKDIR /app
RUN npm config set registry https://registry.npmmirror.com
RUN corepack enable
RUN corepack prepare pnpm@latest --activate
COPY . .
RUN pnpm install
RUN pnpm list tesseract.js || echo "tesseract.js not found (安装失败)"
EXPOSE 8082
CMD ["pnpm", "run", "start"]