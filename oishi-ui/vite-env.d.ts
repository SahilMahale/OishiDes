/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_DEPLOY_ENV: string
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
