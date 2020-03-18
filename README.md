# BUPT-nCov-Hitcarder

北邮nCov肺炎健康打卡脚本

 - python模块


 项目用于学习交流，仅用于各项无异常时打卡，如有身体不适等情况还请自行如实打卡~


## Usage

### 本地运行

1. clone本项目并cd到本目录
    ```bash
    $ git clone https://github.com/timfaner/BUPT-nCov-Hitcarder.git --depth 1
    $ cd BUPT-nCov-Hitcarder
    ```
    
2. 安装依赖

    ```bash
    $ pip3 install -r requirements.txt
    ```

3. 运行一次打卡
    ```bash
    python3 daka.py username password
    ```

4. 定时自动打卡脚本
    - Linux：  使用 `cron`
    - Macos：  使用 `Automator` 或 `cron`
    - Windows：使用 `任务计划程序`


## Issues
如有任何需求、使用遇到问题，请开一个issue

## Thanks

感谢浙大原始代码 [ZJU-nCov-Hitcarder](https://github.com/Tishacy/ZJU-nCov-Hitcarder)


## LICENSE

Copyright (c) 2020 chaunhewie.

Licensed under the [MIT License](https://github.com/chaunhewie/BUPT-nCov-Hitcarder/blob/master/LICENSE)

