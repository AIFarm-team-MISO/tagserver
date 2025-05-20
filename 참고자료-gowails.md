# 🔍 Wails 프로젝트 구성 요소 설명

Wails 프로젝트를 실행하고 개발하기 위해 설치한 주요 구성 요소들의 역할과 필요성을 정리한 문서입니다.

---

## 1. 🐹 Go

### ❓ 무엇인가요?
- 구글이 개발한 오픈소스 프로그래밍 언어로, Wails의 **백엔드 로직**을 담당합니다.
- Wails 앱의 핵심 로직, 데이터 처리, 외부 연동 등을 Go로 작성합니다.

### ✅ 왜 필요한가요?
- Wails는 **Go 언어 기반의 프레임워크**입니다.
- 앱 실행, 빌드, 백엔드 API 구현 등 모든 Wails의 핵심 기능은 Go로 이루어집니다.

---

## 2. 🧵 Wails CLI

### ❓ 무엇인가요?
- Go로 만든 **데스크탑 앱 프레임워크**로, Electron 없이 프론트엔드(UI)와 Go 백엔드를 연결해주는 역할을 합니다.
- `wails init`, `wails dev`, `wails build` 등 명령어 제공

### ✅ 왜 필요한가요?
- Svelte(또는 React 등)로 만든 UI를 Go 백엔드와 연결해 하나의 데스크탑 앱으로 만듭니다.
- 개발 중에는 핫리로드 및 디버깅을 도와줍니다.

---

## 3. 🟩 Node.js

### ❓ 무엇인가요?
- 자바스크립트 런타임으로, Wails 프로젝트의 **프론트엔드 빌드 시스템**을 실행하는 데 사용됩니다.

### ✅ 왜 필요한가요?
- Svelte, Vue, React 같은 프론트엔드 프레임워크는 Node.js 환경에서 컴파일 및 실행됩니다.
- Wails의 `frontend/` 폴더 내부는 Node.js 프로젝트이므로, 빌드를 위해 Node.js가 필요합니다.

---

## 4. 📦 npm (Node Package Manager)

### ❓ 무엇인가요?
- Node.js의 패키지 관리 도구로, Wails 프론트엔드에서 필요한 라이브러리를 설치하고 관리합니다.

### ✅ 왜 필요한가요?
- Wails 프론트엔드에서 사용하는 Svelte 관련 모듈, Vite, 기타 JS 라이브러리를 설치/관리합니다.
- Wails 실행 시 `npm install` 명령이 자동 실행됩니다 (`wails dev` 또는 `wails build`).

---

## 🧩 전체 연동 구조 요약
[ Svelte (frontend) ] ← Node.js + npm
↓
[ Wails Framework ] ← Go
↓
[ Python 연동 (외부 API or IPC) ]


- Node.js + npm → 프론트 빌드 및 실행
- Wails + Go → 백엔드 로직 및 통신
- Python → Wails 앱 호출 및 결과 수신 (api 형식으로 할지는 생각중.. )


✅ 각 파일별 의미와 템플릿과의 독립성
파일	역할	템플릿과의 관계
go.mod	Go 모듈 선언 및 의존성 관리	🔗 프론트엔드 종류와 무관
go.sum	Go 의존성 해시값 (자동 생성됨)	🔗 프론트엔드 종류와 무관
main.go	앱 시작점 (Wails 실행 구조)	🔗 프론트와 연결만 하므로 내부 UI 구조와 무관
app.go (있다면)	Go 백엔드 API 정의	🔗 프론트와 데이터만 주고받음
wails.json	Wails 프로젝트 정보 (타이틀, 번들 등)	🔗 템플릿 종류는 기록되지 않음
README.md	설명 파일	🔗 그냥 문서, 자유롭게 수정 가능