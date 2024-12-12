# minik8s
参考k8s版本1.31.3
mini版的k8s，用于学习k8s设计思想

# kubelet工作原理
[11.深入k8s：kubelet工作原理及其初始化源码分析](https://cloud.tencent.com/developer/article/1701500)

# 个人最初猜想
1. 最开始我以为是服务端制定标准，客户端去实现，实际k8s却是kubelet作为客户端制定了CRI标准,容器厂商作为服务端去实现标准（实现了可扩展性）
2. 最开始我以为是通过shell脚本去派发docker run命令，实际k8s却是通过rpc调用容器接口
3. 最开始我以为kubectl apply -f pod.yaml是kubectl请求apiserver，apiserver通过http请求kubelet，实际k8s是通过witch机制实现
4. 最开始我以为kubelet的watch机制是通过轮询，实际k8s是通过回调实现