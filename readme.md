# spider 🕷️

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## 📄 许可与说明

本项目使用 [GNU General Public License v3.0](LICENSE) 发布。  
您可以自由地：使用、复制、修改、再发布，包括商业用途。  
项目初衷为技术学习与交流；请遵守 GPL-3.0 全部条款。

## 使用

1. 打开spider.exe
2. 输入要搜索的关键词然后回车
3. 输入下载数量(100~300之间)
4. 等待下载完成

## 环境要求

- 网络能正常访问```cn.bing.com``` 和 ```*.cn.bing.net```

## 报错解释

暂无

## 特性

- 记输入的下载数量为 `n`  
  - 若 `n` 不是数字 → 默认下载 150 张  
  - 若 `n < 100` → 按 100 张下载  
  - 若 `n > 300` → 按 300 张下载
- 多线程下载
