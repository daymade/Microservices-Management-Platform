# 使用官方 Node.js 镜像作为基础镜像
FROM node:18-alpine

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY catalog-service-management-ui/package*.json ./

# 安装依赖
RUN npm install

# 复制所有项目文件
COPY catalog-service-management-ui .

# 构建项目
RUN npm run build

# 安装 serve 包用于提供静态文件
RUN npm install -g serve

# 暴露端口
EXPOSE 3000

# 启动 serve 提供静态文件
CMD ["serve", "-s", "dist", "-l", "5173"]
