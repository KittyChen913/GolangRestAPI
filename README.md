# Golang REST API + ELK 日誌系統

這是一個展示如何使用 Go 語言建立 RESTful API，整合 Microsoft SQL Server 資料庫以及 ELK Stack (Filebeat, Elasticsearch, Kibana) 進行集中式日誌管理的範例專案。

所有服務都透過 Docker Compose 管理，方便快速啟動環境。

## 系統架構

應用程式由以下服務組成：

*   **`api-service`**: 核心的 REST API，使用 Go 語言編寫，Log 處理採用具備高效能、支援結構化輸出的 [zap](https://github.com/uber-go/zap) 日誌框架。
*   **`mssql-db`**: Microsoft SQL Server 2022 資料庫。
*   **`filebeat`**: 一個輕量級的日誌收集器，收集日誌並轉發到 Elasticsearch。
*   **`elasticsearch`**: 一個搜尋與分析引擎，用於儲存和索引日誌，此專案用於接收 filebeat 傳過來的日誌內容。
*   **`kibana`**: 一個資料視覺化工具，用於查詢和檢視儲存在 Elasticsearch 中的日誌。

<p align="center">
    <img src="images/GolangRestAPI%20Architecture%20Diagram.png" alt="架構圖" width="75%">
</p>

## 環境需求

-   Docker
-   Docker Compose (通常會跟 Docker 一起安裝)

## 快速開始

### 1. 啟動所有服務

```sh
docker-compose up --build
```

### 2. 服務說明

服務啟動後，您可以透過以下位址存取：

-   **API 服務**: [http://localhost:8080](http://localhost:8080)
-   **Swagger 文件**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
-   **elasticsearch**: [http://localhost:9200](http://localhost:9200)
    -   可訪問 [http://localhost:9200/_cat/indices?v](http://localhost:9200/_cat/indices?v) 查看目前有哪些索引
    -   API 服務被訪問後，filebeat 會自動傳遞 `filebeat-datastream-*` index 過來
-   **Kibana (日誌儀表板)**: [http://localhost:5601](http://localhost:5601)
    -   開啟後，點擊左側選單的 "Discover"，您應該能看到由 `api-service` 產生的日誌。
    -   點選左側選單的 "Stack Management" > "Index Management"，勾選 "Include hidden indices" 即可看到 `.ds-filebeat-datastream-9.0.3-*`
-   **資料庫連線**:
    -   **工具**: SQL Server Management Studio
    -   **主機**: `localhost`
    -   **Port**: `1433`
    -   **使用者**: `SA`
    -   **密碼**: `StrongPassword!123`

## 組態設定

`docker-compose.yml`定義所有服務的啟動方式、相依性、網路與資料卷。

### 環境變數

*   **`api-service`**:
    *   `DemoDb`: 連接到 MS SQL 資料庫的連線字串。
*   **`mssql-db`**:
    *   `SA_PASSWORD`: `SA` 使用者的密碼。

### 資料庫初始化

當 `mssql-db` 服務首次啟動時，會自動執行 `initDb/init.sql` 指令碼。您可以將建立資料表 (`CREATE TABLE`)、插入初始資料 (`INSERT`) 等 SQL 命令放在此檔案中。

### 日誌收集

日誌收集的設定位於 `filebeat/filebeat.yml`，filebeat 會監控 `log-data` 磁碟區中的日誌檔案，該磁碟區與 `api-service` 共享。

## 目錄結構

```
├── api-service/          # Golang API
│   ├── customerrors      # 自定義錯誤
│   ├── db/               # 資料庫連線與操作
│   ├── docs/             # Swagger 文件
│   ├── logger/           # 日誌工具
│   ├── middlewares/      # Gin 中介層
│   ├── models/           # 資料模型
│   ├── routes/           # API 路由
│   ├── utils/            # 工具
│   ├── Dockerfile        # Docker 構建檔案
│   └── main.go
├── filebeat/filebeat.yml # Filebeat 設定
├── initDb/init.sql       # SQL Server 初始化腳本
├── docker-compose.yml    # Docker Compose 配置
```
