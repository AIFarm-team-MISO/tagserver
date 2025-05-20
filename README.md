# 🏷️ 태그서버 프로그램 전체 설계 개요

이 문서는 Wails 기반으로 네이버 스마트스토어 상세페이지의 검색 태그를 수집하는 프로그램을 만드는 전체 구조를 정의합니다.

## ✅ 전체 개발 계획

1. **1단계: Go + Wails를 이용한 태그 수집 기능 구현**

   * 스마트스토어 상세페이지를 Wails 앱의 웹뷰로 열고 검색 태그를 추출
   * 프론트엔드는 Svelte, 백엔드는 Go로 구현

2. **2단계: 수집 기능을 외부에서 접근 가능한 API로 확장**

   * Python 등 외부 프로그램이 HTTP 요청을 통해 Wails 앱에 태그 수집 요청 가능
   * Wails 앱이 실행 중일 때만 외부 요청에 응답

---

## 📐 전체 아키텍처 (2단계 기준)

```plaintext
[Python 코드]
    │  (HTTP 요청: /extract_tags?url=...)
    ▼
[Wails + Svelte 앱 - 태그서버]
    │
    ├─ 📥 Svelte에서 WebView로 스마트스토어 상세페이지 열기
    ├─ 🧠 JavaScript (MutationObserver 등)로 태그 DOM 추출
    ├─ 📤 추출된 태그를 Go로 전달 (`backend.ReceiveTags(tags)`)
    └─ 📡 Go가 HTTP JSON 응답 처리 (/extract_tags API)
```

---

## 📁 프로젝트 구성 예시

```plaintext
tagserver/
├── main.go               ← Go 엔트리포인트 + 서버 구현
├── app.go                ← 백엔드 구조체 및 태그 수집 메서드
├── wails.json            ← Wails 설정파일
├── frontend/
│   ├── App.svelte        ← 페이지 로딩 + 태그 추출 로직 포함
│   └── components/       ← 필요 시 컴포넌트 분리
└── .gitattributes
```

---

## 🔧 주요 기능 설계

### 1. 내부 태그 수집 흐름 (1단계 목표)

* App.svelte가 네이버 상품 상세페이지 로드
* MutationObserver로 DOM 완성 대기
* `ul._3Vox1DKZiA > li > a` 에서 검색 태그 추출
* 추출된 태그를 `window.backend.App.ReceiveTags(tags)`로 전달
* Go 구조체(App)에 태그 저장 및 출력

### 2. Python → Wails 요청 흐름 (2단계 확장)

```http
GET /extract_tags?url=https://smartstore.naver.com/... 
```

* Go에서 이 요청을 처리 → Svelte에 URL 전달 → 태그 추출 완료 시 응답

### 3. 태그 추출 로직 (Svelte 내부)

```javascript
const observer = new MutationObserver(() => {
    const tags = [...document.querySelectorAll("ul._3Vox1DKZiA > li > a")].map(el => el.textContent);
    if (tags.length) {
        window.backend.App.ReceiveTags(tags);
        observer.disconnect();
    }
});
```

### 4. Go로 결과 전달

```go
func (a *App) ReceiveTags(tags []string) {
    a.lastResult = tags
}
```

### 5. JSON API 응답 처리 (2단계에서 적용)

```go
func (a *App) ExtractTags(url string) []string {
    a.lastResult = nil
    a.currentURL = url
    // Svelte에서 자동으로 열고 추출되면 lastResult에 저장됨
    waitForTags()
    return a.lastResult
}
```

---

## 📡 결과 예시 (Python 수신)

```json
{
  "tags": ["캠핑", "차박", "휴대용"]
}
```

---

## 🧪 테스트 계획

* 네이버 스마트스토어 제품 3개 이상으로 테스트
* 태그 없는 페이지 처리 예외 확인
* Python → Wails 동기화 타이밍 확인

---

## 🛠️ 향후 확장 가능성

* 여러 URL 일괄 처리
* 브라우저 세션 유지 (로그인 필요 상품 처리)
* Python 쪽에서 대량 호출 자동화


# 프로그램 개발 시작 