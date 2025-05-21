export {};

declare global {
  interface Window {
    backend: {
      Crawler: {
        FetchTags(keyword: string): Promise<string[]>;
      };
    };
  }
}
