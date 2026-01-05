/// <reference types="vite/client" />
export {};

declare global {
  interface Window {
    appConfig: {
      basePath: string;
    };
  }
}