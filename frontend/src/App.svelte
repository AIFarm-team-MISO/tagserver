<script>
  import { onMount } from 'svelte';

  let targetUrl = "https://shopping.naver.com/ns/home";

  onMount(() => {
    window.location.href = targetUrl;
  });
</script>

<main>
  <h1>🔗 네이버 쇼핑 페이지에 진입 중...</h1>
</main>
