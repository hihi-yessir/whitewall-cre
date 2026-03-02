# ============================================
# Stage 1: Build the Go server
# ============================================
FROM golang:1.25.5-bookworm AS builder

#작업 디렉토리 /app으로 함
WORKDIR /app

# Go 모듈 파일 복사 및 다운로드
COPY cre-wrapper/go.mod cre-wrapper/go.sum ./cre-wrapper/
RUN cd cre-wrapper && go mod download

# 소스 코드 복사
COPY cre-wrapper/ ./cre-wrapper/

# Go 서버 빌드
RUN cd cre-wrapper && CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# ============================================
# Stage 2: Runtime with CRE CLI
# ============================================
FROM ubuntu:24.04

# 필수 패키지 설치 (GLIBC 2.39 포함)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    git \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Go 설치 (CRE 워크플로우 컴파일에 필요)
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then GOARCH="amd64"; else GOARCH="arm64"; fi && \
    wget -q https://go.dev/dl/go1.25.5.linux-${GOARCH}.tar.gz && \
    tar -C /usr/local -xzf go1.25.5.linux-${GOARCH}.tar.gz && \
    rm go1.25.5.linux-${GOARCH}.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# CRE CLI 설치
RUN curl -sSL https://cre.chain.link/install.sh | bash
ENV PATH="/root/.cre/bin:${PATH}"

WORKDIR /app

# 빌드된 서버 복사
COPY --from=builder /app/server ./server

# CRE 워크플로우 및 Go 모듈 복사
COPY go.mod go.sum ./
COPY whitewall-access/ ./whitewall-access/
COPY contracts/ ./contracts/

# Go 모듈 미리 다운로드 (런타임에 다운로드 안 하도록)
RUN go mod download

# 환경변수 설정
ENV PORT=8080
ENV CRE_PROJECT_ROOT=/app
ENV CRE_ETH_PRIVATE_KEY=""
ENV CRE_CREDENTIALS=""
ENV CRE_PROJECT_YAML=""
ENV CRE_TARGET="staging-settings"

# 포트 노출구
EXPOSE 8080

# 시작 스크립트 생성 (CRE credentials + project.yaml 설정 + 워크플로우 사전 컴파일 후 서버 실행)
RUN echo '#!/bin/bash\n\
set -e\n\
\n\
# 1. CRE credentials 설정\n\
if [ -n "$CRE_CREDENTIALS" ]; then\n\
  mkdir -p /root/.cre\n\
  echo "$CRE_CREDENTIALS" | base64 -d > /root/.cre/cre.yaml\n\
  echo "CRE credentials configured"\n\
fi\n\
\n\
# 2. project.yaml 설정\n\
if [ -n "$CRE_PROJECT_YAML" ]; then\n\
  echo "$CRE_PROJECT_YAML" | base64 -d > /app/project.yaml\n\
  echo "CRE project.yaml configured"\n\
fi\n\
\n\
# 3. CRE 워크플로우 사전 컴파일 (서버 시작 전)\n\
echo "Pre-compiling CRE workflow..."\n\
cd /app/whitewall-access\n\
if cre workflow build . --target ${CRE_TARGET:-staging-settings} 2>&1; then\n\
  echo "CRE workflow pre-compiled successfully"\n\
else\n\
  echo "CRE workflow build command not available, will compile on first request"\n\
fi\n\
\n\
# 4. 서버 실행\n\
cd /app\n\
exec ./server' > /app/start.sh && chmod +x /app/start.sh

# 서버 실행
CMD ["/app/start.sh"]
