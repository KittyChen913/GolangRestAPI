# Golang REST API + ELK 日誌系統

這是一個展示如何使用 Go 語言建立 RESTful API，整合 Microsoft SQL Server 資料庫以及 ELK Stack (Filebeat, Elasticsearch, Kibana) 進行集中式日誌管理的範例專案。

所有服務都透過 Docker Compose 管理，方便快速啟動環境。

## 系統架構

應用程式由以下服務組成：

*   **`api-service`**: 核心的 REST API，使用 Go 語言與 [Gin](https://github.com/gin-gonic/gin) 框架編寫，Log 處理採用具備高效能、支援結構化輸出的 [zap](https://github.com/uber-go/zap) 日誌框架。
*   **`mssql-db`**: Microsoft SQL Server 2022 資料庫。
*   **`filebeat`**: 一個輕量級的日誌收集器，收集日誌並轉發到 Elasticsearch。
*   **`elasticsearch`**: 一個搜尋與分析引擎，用於儲存和索引日誌，此專案用於接收 filebeat 傳過來的日誌內容。
*   **`kibana`**: 一個資料視覺化工具，用於查詢和檢視儲存在 Elasticsearch 中的日誌。

<p align="center">
    <img src="images/GolangRestAPI%20Architecture%20Diagram.png" alt="架構圖" width="75%">
</p>

## 依賴注入（Dependency Injection）

這個專案採用依賴注入設計，Service/Repository 皆於 `main.go` 統一建立並注入到 handler 及 middleware，讓程式碼更易於測試、維護與擴充。

```go
// main.go
userRepo := repositories.NewUserRepository(db.Db)
userService := services.NewUserService(userRepo)
userHandler := handlers.NewUserHandler(userService)

server.POST("/CreateUser", userHandler.CreateUser)
```

## 組態設定

`docker-compose.yml`定義所有服務的啟動方式、相依性、資料卷。

### 環境變數

*   **`api-service`**:
    *   `DemoDb`: 連接到 MS SQL 資料庫的連線字串。
*   **`mssql-db`**:
    *   `SA_PASSWORD`: `SA` 使用者的密碼。

### 資料庫初始化

當 `mssql-db` 服務首次啟動時，會自動執行 `initDb/init.sql` 指令碼，這裡會執行建立資料表 (`CREATE TABLE`) 的 SQL 命令。

### 日誌收集

日誌收集的設定位於 `filebeat/filebeat.yml`，filebeat 會監控 `log-data` 磁碟區中的日誌檔案，該磁碟區與 `api-service` 共享。


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
    -   點選左側選單的 "Stack Management" > "Index Management"，勾選 "Include hidden indices" 即可看到 `.ds-filebeat-datastream-9.0.3-*`，不用自己手動 Create Index
-   **資料庫連線**:
    -   **工具**: SQL Server Management Studio
    -   **主機**: `localhost`
    -   **Port**: `1433`
    -   **使用者**: `SA`
    -   **密碼**: `StrongPassword!123`


## 目錄結構

```
├── api-service/                # Golang API 主程式
│   ├── customerrors/           # 自定義錯誤型別
│   ├── db/                     # 資料庫連線與操作
│   ├── docs/                   # Swagger 自動產生的 API 文檔
│   ├── handlers/               # HTTP handler
│   ├── logger/                 # 日誌工具
│   ├── middlewares/            # Gin 中介層（權限驗證、日誌、錯誤處理）
│   ├── models/                 # 資料模型
│   ├── repositories/           # 資料存取層
│   ├── routes/                 # API 路由註冊
│   ├── services/               # 商業邏輯層
│   ├── utils/                  # 工具函式
│   ├── Dockerfile              # Docker 構建檔案
│   ├── go.mod                  # Go module 設定
│   ├── go.sum                  # Go module 依賴雜湊
│   └── main.go                 # 進入點
├── filebeat/
│   └── filebeat.yml            # Filebeat 設定檔
├── images/                     # 說明文件用的圖檔
├── initDb/
│   └── init.sql                # SQL Server 初始化腳本
├── docker-compose.yml          # Docker Compose 配置
└── README.md                   # 專案說明文件
```
