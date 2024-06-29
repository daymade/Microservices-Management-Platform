# 基础镜像，用于安装 migrate 工具
FROM alpine:latest AS migrate-installer
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        MIGRATE_URL="https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz"; \
    elif [ "$ARCH" = "aarch64" ]; then \
        MIGRATE_URL="https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-arm64.tar.gz"; \
    else \
        echo "unsupported architecture: $ARCH"; exit 1; \
    fi && \
    wget -O migrate.tar.gz $MIGRATE_URL && \
    tar -xvf migrate.tar.gz -C /usr/local/bin && \
    chmod +x /usr/local/bin/migrate && \
    rm migrate.tar.gz

# 最终镜像
FROM alpine:latest
WORKDIR /usr/local/bin
COPY --from=migrate-installer /usr/local/bin/migrate /usr/local/bin/migrate

CMD ["migrate"]
