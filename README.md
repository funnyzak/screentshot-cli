# Screenshot-CLI 截图工具

一个轻量级、无依赖的命令行截图工具，支持全屏和区域截图，具有灵活的输出控制和批量处理功能。支持 Windows、macOS 和 Linux 平台。

## 功能特性

- **全屏截图**: 捕获整个屏幕
- **区域截图**: 通过坐标捕获特定区域
- **多种格式**: 支持 PNG、JPG、BMP、GIF 格式
- **质量控制**: 可调节 JPEG 压缩质量
- **剪贴板支持**: 直接将截图复制到剪贴板
- **批量处理**: 可配置间隔时间进行多次截图
- **模板系统**: 支持变量的动态文件名生成
- **跨平台**: 支持 Windows、macOS 和 Linux

## 安装

### 前置要求
- Go 1.21 或更高版本

### 从源码构建
```bash
git clone git@github.com/funnyzak/screenshot-cli.git
cd screenshot-cli
go build -o sshot cmd/main.go
```

### 全局安装
```bash
go install ./cmd/main.go
```

## 使用方法

### 基础截图

```bash
# 全屏截图 (保存为 screenshot.png)
sshot

# 全屏截图并指定文件名
sshot -o desktop.png

# 区域截图
sshot -r "100,100,800,600" -o region.png

# 指定格式和质量
sshot -f jpg -q 80 -o screen.jpg
```

### 输出控制

```bash
# 仅复制到剪贴板
sshot -c

# 使用文件名模板
sshot -t "screenshot_{datetime}.png"

# 多个模板变量
sshot -t "screen_{date}_{counter}.jpg" -f jpg
```

### 批量处理

```bash
# 每3秒拍摄10张截图
sshot -n 10 -i 3 -p "batch" -d "./screenshots"

# 批量区域截图
sshot -r "0,0,1920,1080" -n 5 -i 2 -t "monitor_{counter}.png"

# 批量处理并自定义目录结构
sshot -n 20 -i 1 -d "./captures/{date}" -t "shot_{time}_{counter}.png"
```

## 命令行选项

### 全局选项
| 选项 | 描述 |
|------|------|
| `--help, -h` | 显示帮助信息 |
| `--version, -v` | 显示版本信息 |
| `--verbose` | 启用详细输出 |

### 截图选项
| 选项 | 描述 | 默认值 |
|------|------|--------|
| `--region, -r` | 区域截图 "x,y,width,height" | - |
| `--output, -o` | 输出文件路径 | screenshot.png |
| `--display, -d` | 要捕获的显示器索引 (0为主显示器) | 0 |

### 输出控制
| 选项 | 描述 | 默认值 |
|------|------|--------|
| `--format, -f` | 输出格式 (png/jpg/bmp/gif) | png |
| `--quality, -q` | JPG 压缩质量 (1-100) | 90 |
| `--clipboard, -c` | 复制到剪贴板 | false |
| `--template, -t` | 文件名模板 | - |

### 批量处理
| 选项 | 描述 | 默认值 |
|------|------|--------|
| `--count, -n` | 截图数量 | 1 |
| `--interval, -i` | 截图间隔 (秒) | 1 |
| `--prefix, -p` | 文件名前缀 | shot |
| `--directory` | 输出目录 | . |

## 模板变量

| 变量 | 描述 | 示例 |
|------|------|------|
| `{timestamp}` | Unix 时间戳 | 1640995200 |
| `{datetime}` | 日期和时间 | 20220101_120000 |
| `{date}` | 仅日期 | 20220101 |
| `{time}` | 仅时间 | 120000 |
| `{counter}` | 序列号 | 001, 002, 003 |
| `{random}` | 随机字符串 | a1b2c3 |
| `{prefix}` | 文件名前缀 | shot |

## 详细使用示例

### 1. 基础截图操作

```bash
# 快速截图到当前目录
sshot

# 指定输出路径和格式
sshot -o ~/Desktop/screen.png -f png

# 高质量JPEG截图
sshot -f jpg -q 95 -o high_quality.jpg

# 低质量快速截图
sshot -f jpg -q 30 -o quick.jpg
```

### 2. 区域截图应用

```bash
# 捕获屏幕左上角区域
sshot -r "0,0,800,600" -o top_left.png

# 捕获屏幕中央区域
sshot -r "560,315,800,450" -o center.png

# 捕获任务栏区域 (假设在底部)
sshot -r "0,1040,1920,40" -o taskbar.png

# 捕获浏览器窗口区域
sshot -r "100,100,1200,800" -o browser.png
```

### 3. 剪贴板操作

```bash
# 截图并复制到剪贴板
sshot -c

# 区域截图到剪贴板
sshot -r "0,0,800,600" -c

# 截图到剪贴板并保存文件
sshot -c -o backup.png
```

### 4. 开发工作流

```bash
# 开发过程中定时截图
sshot -n 50 -i 5 -d "./dev-screenshots" -t "dev_{datetime}_{counter}.png"

# 代码审查截图
sshot -r "0,0,1200,800" -n 10 -i 30 -t "review_{date}_{counter}.png"

# 测试自动化截图
sshot -n 100 -i 2 -d "./test-results" -t "test_{timestamp}_{counter}.png"

# 性能监控截图
sshot -n 60 -i 60 -d "./performance" -t "perf_{time}.png"
```

### 5. 文档制作

```bash
# 创建教程截图
sshot -r "0,0,1200,800" -f jpg -q 85 -t "tutorial_{date}_{counter}.jpg"

# 软件界面截图
sshot -r "100,100,1000,700" -f png -t "ui_{date}_{counter}.png"

# 错误报告截图
sshot -r "0,0,1920,1080" -f jpg -q 90 -t "error_{datetime}.jpg"
```

### 6. 监控和记录

```bash
# 每小时截图一次，持续24小时
sshot -n 24 -i 3600 -d "./daily-monitor" -t "day_{date}_{time}.png"

# 每5分钟截图一次，持续一周
sshot -n 2016 -i 300 -d "./weekly-monitor" -t "week_{datetime}.png"

# 系统状态监控
sshot -r "0,0,800,600" -n 1440 -i 60 -d "./system-monitor" -t "sys_{time}.png"
```

### 7. 创意和设计

```bash
# 设计灵感收集
sshot -r "0,0,1920,1080" -n 50 -i 120 -d "./inspiration" -t "inspire_{random}.png"

# 颜色提取截图
sshot -r "100,100,200,200" -f png -t "color_{timestamp}.png"

# 布局分析截图
sshot -r "0,0,1200,800" -f jpg -q 80 -t "layout_{date}_{counter}.jpg"
```

### 8. 教育和培训

```bash
# 课程录制截图
sshot -r "0,0,1920,1080" -n 1800 -i 2 -d "./course" -t "lesson_{counter}.png"

# 演示步骤截图
sshot -r "0,0,1200,800" -n 20 -i 10 -t "demo_step_{counter}.png"

# 学生作业截图
sshot -r "0,0,800,600" -f jpg -q 85 -t "homework_{date}_{random}.jpg"
```

### 9. 游戏和娱乐

```bash
# 游戏精彩时刻截图
sshot -r "0,0,1920,1080" -f png -t "game_{timestamp}.png"

# 视频帧提取
sshot -r "0,0,1280,720" -n 300 -i 1 -d "./frames" -t "frame_{counter:04d}.png"

# 直播截图
sshot -r "0,0,1920,1080" -n 1000 -i 5 -d "./stream" -t "stream_{time}.png"
```

### 10. 高级模板使用

```bash
# 使用复杂模板
sshot -t "screenshot_{date}_{time}_{random}_{counter:03d}.png"

# 按日期组织文件
sshot -d "./screenshots/{date}" -t "{time}_{counter}.png"

# 按项目组织文件
sshot -d "./projects/{project_name}/{date}" -t "{time}_{description}.png"

# 使用环境变量
sshot -t "screenshot_{USER}_{date}_{time}.png"
```

### 11. 多显示器支持

```bash
# 主显示器截图
sshot -d 0 -o primary.png

# 第二显示器截图
sshot -d 1 -o secondary.png

# 所有显示器批量截图
sshot -d 0 -t "display0_{counter}.png" -n 5 &
sshot -d 1 -t "display1_{counter}.png" -n 5 &
```

### 12. 自动化脚本集成

```bash
#!/bin/bash
# 自动化截图脚本示例

# 创建目录
mkdir -p "./screenshots/$(date +%Y%m%d)"

# 定时截图
sshot -n 10 -i 30 -d "./screenshots/$(date +%Y%m%d)" -t "auto_{time}.png"

# 条件截图
if [ $? -eq 0 ]; then
    echo "截图成功"
else
    echo "截图失败"
fi
```

## 错误处理

工具提供清晰的错误信息和适当的退出代码：

- `0`: 成功
- `1`: 参数错误
- `2`: 文件系统错误
- `3`: 截图捕获错误
- `4`: 格式转换错误
- `5`: 剪贴板错误

## 性能指标

- **响应时间**: 全屏 < 500ms，区域 < 300ms
- **内存使用**: < 50MB
- **文件大小**: < 10MB 可执行文件
- **CPU 使用**: 捕获期间 < 20%

## 开发

### 项目结构
```
screenshot-cli/
├── cmd/
│   └── main.go                 # 程序入口点
├── internal/
│   ├── capture/
│   │   └── screen.go          # 截图核心逻辑
│   ├── output/
│   │   ├── format.go          # 格式转换
│   │   ├── file.go            # 文件输出
│   │   └── clipboard.go       # 剪贴板操作
│   ├── batch/
│   │   └── processor.go       # 批量处理逻辑
│   └── config/
│       ├── args.go            # 命令行参数解析
│       └── template.go        # 文件名模板处理
├── go.mod
└── README.md
```

### 构建

```bash
# 开发构建
go build -o sshot cmd/main.go

# 发布构建 (Windows)
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o sshot.exe cmd/main.go

# 发布构建 (macOS)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o sshot cmd/main.go

# 发布构建 (Linux)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o sshot cmd/main.go
```

### 测试

```bash
# 运行所有测试
go test ./...

# 运行覆盖率测试
go test -cover ./...

# 运行特定测试
go test ./internal/capture
```

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 LICENSE 文件。

## 贡献

1. Fork 本仓库
2. 创建功能分支
3. 进行更改
4. 为新功能添加测试
5. 确保所有测试通过
6. 提交拉取请求

## 致谢

- [kbinani/screenshot](https://github.com/kbinani/screenshot) - 跨平台截图库
- [spf13/cobra](https://github.com/spf13/cobra) - 命令行框架
- [atotto/clipboard](https://github.com/atotto/clipboard) - 跨平台剪贴板库 