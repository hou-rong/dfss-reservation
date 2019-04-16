# dfss-reservation

## 简介

东方时尚自动约车工具。一个简单的脚本，主要解决想约周六日的课程，但 7 点起不来的问题。

## 实现及流程
### 实现
- API
    - login 鉴权
    - get_plain 可约课程列表，获取最佳约车时间
    - reservation 约车
- CMD
    - main 流程控制
    
### 流程

1. 模拟登录，获取鉴权信息和校验和
2. 请求约车列表接口，计算出最佳约车时间
3. 携带最佳约车时间请求约车接口完成约车

## 配置

- USERNAME 你的学号（小黄本上的数字）
- PASSWORD 你的密码
- TrainingTimeSlotScore 不同时间段的分数（越高在有课的情况下越先选到）
