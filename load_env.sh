#!/bin/bash

echo "Starting to load environment variables from .env file..."

# 确保 .env 文件存在
if [ ! -f ".env" ]; then
    echo ".env file not found!"
    exit 1
fi

# 从 .env 文件加载环境变量
while IFS='=' read -r key value
do
  if [ -n "$key" ]; then
    # 删除可能的前后空格
    key=$(echo $key | xargs)
    value=$(echo $value | xargs)

    # 使用 eval 来处理值中包含的特殊字符
    eval export $key='$value'
    echo "Exported: $key = '${!key}'"  # 使用间接引用来打印环境变量的值
  fi
done < .env

echo "All environment variables loaded and printed successfully."
# source ./load_env.sh 

