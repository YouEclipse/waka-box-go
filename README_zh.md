# 


<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/4658208/60469862-2e40bf00-9c2c-11e9-87f7-afe164648de4.png">
  <h3 align="center">waka-box-go</h3>
  <p align="center">📊 一个能够自动更新的 waketime 每周统计的 <a href="https://gist.github.com/">Github Gist</a>. </p>
  <p align="center">  Golang 实现，查看原始版本  <a href="https://github.com/matchai/waka-box">waka-box</a>
  <p align="center">
    <a href="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master"><img src="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master" alt="Update gist with WakaTime stats"></a>
  </p>
</p>

---
[English](./README.md) | 简体中文



> 📌✨ 查看更多像这样的 Pinned Gist 项目,传送门:  https://github.com/matchai/awesome-pinned-gists



## 💻 安装

### 🎒 前置工作

1. 创建一个公开的 GitHub Gist (https://gist.github.com/)
1. 创建一个拥有 `gist` 权限的 token 并复制. (https://github.com/settings/tokens/new)
1. 创建一个 WakaTime 账号(如果已经有了可以跳过) (https://wakatime.com/signup)
1. 在 WakaTime 的 profile settings (https://wakatime.com/settings/profile) 确保 `Display coding activity publicly` 和 `Display languages, editors, operating systems publicly` 被勾选了
1. 在你的 WakaTime 的 account settings, 复制 WakaTime API Key (https://wakatime.com/settings/api-key)

### 🚀 开始安装

1. Fork 这个仓库
1. 编辑  `.github/workflows/schedule.yml` 中的[环境变量](https://github.com/YouEclipse/waka-box-go/actions/runs/126970182/workflow#L17-L19) :

   - **GIST_ID:** ID 是 gist url 的后缀 : `https://gist.github.com/YouEclipse/`**`9bc7025496e478f439b9cd43eba989a4`**.

1. 前往 fork 后的仓库的 **Settings > Secrets**
1. 添加以下环境变量:
   - **GH_TOKEN:** 前置工作中生成的 github token.
   - **WAKATIME_API_KEY:** WakaTime 的 API key. 
  
## 🕵️ 工作原理
- 基于 WakaTime API 获取统计数据
- 基于 Github API 更新 Gist
- 使用 Github Actions 自动更新 Gist  

## 📄  开源协议
本项目使用 [Apache-2.0](./LICENSE) 协议