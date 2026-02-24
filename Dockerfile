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
    && rm -rf /var/lib/apt/lists/*

# CRE CLI 설치
RUN curl -sSL https://cre.chain.link/install.sh | bash
ENV PATH="/root/.cre/bin:${PATH}"

WORKDIR /app

# 빌드된 서버 복사
COPY --from=builder /app/server ./server

# CRE 워크플로우 복사
COPY whitewall-access/ ./whitewall-access/

# 환경변수 설정
ENV PORT=8080
ENV CRE_PROJECT_ROOT=/app
ENV CRE_ETH_PRIVATE_KEY=""
ENV CRE_CREDENTIALS=""
ENV CRE_TARGET="staging-settings"

# 포트 노출구
EXPOSE 8080

# 시작 스크립트 생성 (CRE credentials 설정 후 서버 실행)
# CRE_CREDENTIALS는 base64로 인코딩된 cre.yaml 내용
RUN echo '#!/bin/bash\n\
if [ -n "$CRE_CREDENTIALS" ]; then\n\
  mkdir -p /root/.cre\n\
  echo "$CRE_CREDENTIALS" | base64 -d > /root/.cre/cre.yaml\n\
  echo "CRE credentials configured"\n\
fi\n\
exec ./server' > /app/start.sh && chmod +x /app/start.sh

# 서버 실행
CMD ["/app/start.sh"]
