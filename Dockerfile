# 使用官方 Node.js 基础镜像
FROM node:16.17.1-alpine

# 设置工作目录
WORKDIR /usr/src/app

# 安装依赖项
RUN npm config set registry https://registry.npm.taobao.org/

# 将 package.json  复制到工作目录
COPY ./vue-blog-admin/package.json ./vue-blog-admin/
# 将项目代码复制到工作目录
COPY ./vue-blog-admin ./vue-blog-admin/

# 将 package.json  复制到工作目录
COPY ./vue-blog-front/package.json ./vue-blog-front/
# 将项目代码复制到工作目录
COPY ./vue-blog-front ./vue-blog-front/

#构建前台
WORKDIR /usr/src/app/vue-blog-front/
RUN npm config set registry https://registry.npm.taobao.org/
RUN npm install 
RUN yarn run build 

#构建后台
WORKDIR /usr/src/app/vue-blog-admin/
RUN npm config set registry https://registry.npm.taobao.org/
RUN npm install 
RUN yarn run build 

# 设置 Nginx 作为 Web 服务器
FROM nginx:latest

# 拷贝构建好的项目文件到 Nginx 的默认静态文件目录
COPY --from=0 /usr/src/app/vue-blog-front/dist /usr/share/nginx/html/vue-blog-front/
COPY --from=0 /usr/src/app/vue-blog-admin/dist /usr/share/nginx/html/vue-blog-admin/
# 暴露端口
EXPOSE 80

# 容器启动时自动运行 Nginx 服务器
CMD ["nginx", "-g", "daemon off;"]