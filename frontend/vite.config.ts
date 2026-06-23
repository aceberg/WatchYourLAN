import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'
import { copyFileSync, cpSync, mkdirSync, readdirSync, rmSync } from 'node:fs'
import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'

const rootDir = dirname(fileURLToPath(import.meta.url))

function copyVendorAssets() {
  const vendorDir = resolve(rootDir, 'dist/vendor')

  rmSync(vendorDir, { recursive: true, force: true })
  mkdirSync(vendorDir, { recursive: true })

  const themesDir = resolve(rootDir, 'node_modules/aceberg-bootswatch-fork/dist')
  for (const theme of readdirSync(themesDir, { withFileTypes: true })) {
    if (!theme.isDirectory()) {
      continue
    }

    const targetDir = resolve(vendorDir, 'aceberg-bootswatch-fork/dist', theme.name)
    mkdirSync(targetDir, { recursive: true })
    copyFileSync(
      resolve(themesDir, theme.name, 'bootstrap.min.css'),
      resolve(targetDir, 'bootstrap.min.css'),
    )
  }

  mkdirSync(resolve(vendorDir, 'bootstrap-icons/font'), { recursive: true })
  cpSync(
    resolve(rootDir, 'node_modules/bootstrap-icons/font/bootstrap-icons.min.css'),
    resolve(vendorDir, 'bootstrap-icons/font/bootstrap-icons.min.css'),
  )
  cpSync(
    resolve(rootDir, 'node_modules/bootstrap-icons/font/fonts'),
    resolve(vendorDir, 'bootstrap-icons/font/fonts'),
    { recursive: true },
  )
}

export default defineConfig({
  plugins: [
    solid(),
    {
      name: 'copy-vendor-assets',
      closeBundle: copyVendorAssets,
    },
  ],
  build: {
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`
      }
    }
  }
})
