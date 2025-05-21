# 🏷️ 태그서버 프로그램 전체 설계 개요

이 문서는 Wails 기반으로 네이버 스마트스토어 상세페이지의 검색 태그를 수집하는 프로그램을 만드는 전체 구조를 정의합니다.

## ✅ 전체 개발 계획

1. **1단계: Go + Wails를 이용한 태그 수집 기능 구현**

   * UI 모드: Svelte 웹뷰에서 키워드 입력 → 상위 상품 검색 → 검색 태그 추출 및 화면에 출력
   * 이후 API 모드로 확장하여 외부 프로그램(Python 등)과 통신

2. **2단계: 수집 기능을 외부에서 접근 가능한 API로 확장**

   * Python 등 외부 프로그램이 HTTP 요청을 통해 Wails 앱에 태그 수집 요청 가능
   * Wails 앱이 실행 중일 때만 외부 요청에 응답

---

## 📐 전체 아키텍처 (1단계: UI 모드 기준)

```plaintext
[사용자 입력]
    │  (키워드 입력)
    ▼
[Wails + Svelte 앱 - 태그서버 UI 모드]
    │
    ├─ 📥 키워드 입력 → 검색 페이지 로드
    ├─ 🧠 JS: 상위 상품 링크 탐색
    ├─ 🧠 JS: 각 상품 상세 페이지 접근 + 검색 태그 추출
    ├─ 📤 태그 리스트 Go에 전달
    └─ 🖥️ 화면에 태그 출력
```

---

## 📁 프로젝트 구성 예시

```plaintext
tagserver/
├── main.go               ← Go 진입점 + 앱 실행
├── app.go                ← 백엔드 로직: 태그 저장, Svelte 바인딩 함수
├── wails.json            ← Wails 설정파일
├── frontend/
│   ├── App.svelte        ← 키워드 입력 + 태그 출력 UI
│   └── components/       ← 필요 시 분리
└── .gitattributes
```

---

## 🔧 주요 기능 설계 (UI 모드)

### 1. App.svelte 구조 (사용자 입력 → 검색)

```svelte
<script>
  import { onMount } from 'svelte';
  let keyword = "";
  let tags = [];

  async function search() {
    if (!keyword) return;
    tags = await window.backend.App.ExtractTags(keyword);
  }
</script>

<main>
  <h1>🔍 키워드 기반 태그 추출</h1>
  <input placeholder="검색어 입력" bind:value={keyword} on:keydown={(e) => e.key === 'Enter' && search()} />
  <button on:click={search}>검색</button>

  {#if tags.length > 0}
    <ul>
      {#each tags as tag}
        <li>{tag}</li>
      {/each}
    </ul>
  {/if}
</main>
```

### 2. app.go - Go에서 바인딩 함수 구현

```go
func (a *App) ExtractTags(keyword string) []string {
    // 1. 검색 URL 생성
    // 2. HTML 크롤링 (예: 상위 상품 상세페이지 1~3개 진입)
    // 3. ul._3Vox1DKZiA > li > a 태그 추출
    // (현재는 mock 데이터 반환)
    return []string{"태그1", "태그2", "태그3"}
}
```

### 3. main.go - 바인딩 포함 확인

```go
Bind: []interface{}{
    app,
},
```

---

## 🧪 테스트 계획

* 검색어를 입력 후 태그 목록이 뜨는지 UI에서 확인
* 향후 실제 DOM 파싱 / iframe 내 탐색 / MutationObserver 적용 예정

---

## 🛠️ 이후 확장 계획

* API 모드 개발 → Python에서 REST 요청 처리
* 상세페이지 실시간 DOM 접근 시 iframe 우회 처리 또는 headless 브라우저 연동





# 프로그램 개발 시작 
