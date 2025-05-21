<script>
// @ts-nocheck

  let keyword = '';
  let tags = [];
  let loading = false;

  async function fetchTags() {
    loading = true;
    try {
      const result = await window.backend.Crawler.FetchTags(keyword);
      tags = result;
      console.log("✅ 받은 태그:", result);
    } catch (e) {
      console.error("❌ 태그 요청 실패:", e);
    } finally {
      loading = false;
    }
  }
</script>

<main>
  <h1>상품 태그 추출기</h1>

  <input bind:value={keyword} placeholder="검색어를 입력하세요" />
  <button on:click={fetchTags} disabled={loading}>검색</button>

  {#if loading}
    <p>로딩 중...</p>
  {:else if tags.length > 0}
    <ul>
      {#each tags as tag}
        <li>{tag}</li>
      {/each}
    </ul>
  {:else}
    <p>태그 없음</p>
  {/if}
</main>
