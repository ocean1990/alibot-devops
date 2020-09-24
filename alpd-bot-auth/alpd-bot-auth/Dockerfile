FROM registry.cn-hangzhou.aliyuncs.com/alpd/node:14.9.0-alpine

COPY package.json package-lock.json *.js /app/
COPY protos/auth.proto /app/protos/
RUN npm config set registry https://registry.npm.taobao.org
RUN cd /app && npm install
WORKDIR /app

EXPOSE 9001

ENTRYPOINT ["node", "index.js"]
