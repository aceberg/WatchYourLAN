/* @refresh reload */
import { render } from 'solid-js/web'
import '@fontsource/open-sans/latin-400.css'
import '@fontsource/open-sans/latin-700.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
import App from './App.tsx'
import { applyBootColorMode } from './functions/theme.ts'

applyBootColorMode();

const root = document.getElementById('root')

render(() => <App />, root!)
