## Usage

```bash
$ npm install # or pnpm install or yarn install
```

## Local Web Preview

Start the safe mock API in one terminal:

```bash
npm run mock:api
```

Start Vite in another terminal:

```bash
npm run dev -- --host 127.0.0.1
```

Open [http://127.0.0.1:5173](http://127.0.0.1:5173). Vite proxies `/api` and `/fs` to the mock API at `http://127.0.0.1:8840`. The mock API returns static development data only; it does not scan the LAN, send notifications, send Wake-on-LAN packets, run port scans, or write persistent data.

### Learn more on the [Solid Website](https://solidjs.com) and come chat with us on our [Discord](https://discord.com/invite/solidjs)

## Available Scripts

In the project directory, you can run:

### `npm run dev`

Runs the app in the development mode.<br>
Open [http://localhost:5173](http://localhost:5173) to view it in the browser.

### `npm run dev:local`

Runs the app in development mode on [http://127.0.0.1:5173](http://127.0.0.1:5173).

### `npm run mock:api`

Runs the development-only mock API on [http://127.0.0.1:8840](http://127.0.0.1:8840).

### `npm run build`

Builds the app for production to the `dist` folder.<br>
It correctly bundles Solid in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

## Deployment

Learn more about deploying your application with the [documentations](https://vite.dev/guide/static-deploy.html)
